package locker

import (
	"context"
	"fmt"
	"strings"

	"github.com/Scalingo/go-utils/logger"
	"github.com/Scalingo/link/models"
	"github.com/coreos/etcd/clientv3"
	"github.com/pkg/errors"
)

type etcdLocker struct {
	etcd      *clientv3.Client
	leaseID   clientv3.LeaseID
	leaseTime int64
	key       string
}

func NewETCDLocker(etcd *clientv3.Client, ip string) *etcdLocker {
	key := fmt.Sprintf("%s/default/%s", models.ETCD_LINK_DIRECTORY, strings.Replace(ip, "/", "_", -1))
	return &etcdLocker{
		etcd:      etcd,
		key:       key,
		leaseTime: 5,
	}
}

func (l *etcdLocker) Refresh(ctx context.Context) error {
	log := logger.Get(ctx)

	if l.leaseID == 0 {
		grant, err := l.etcd.Grant(ctx, l.leaseTime)
		if err != nil {
			return errors.Wrap(err, "fail to generate grant")
		}

		l.leaseID = grant.ID
	}

	// The goal of this transaction is to create the key with our leaseID only if this key does not exist
	// We use a transaction to make sure that concurrent tries wont interfere with each others.

	_, err := l.etcd.Txn(ctx).
		// If the key does not exists (createRevision == 0)
		If(clientv3.Compare(clientv3.CreateRevision(l.key), "=", 0)).
		// Create it with our leaseID
		Then(clientv3.OpPut(l.key, "locked", clientv3.WithLease(l.leaseID))).
		Commit()
	if err != nil {
		return errors.Wrap(err, "fail to refresh lock")
	}

	_, err = l.etcd.KeepAliveOnce(ctx, l.leaseID)
	if err != nil {
		// We got an error while sending keepalive: Regenerate lease
		l.leaseID = 0
		log.WithError(err).Error("Keep alive failed, resetting lease")
	}

	return nil
}

func (l *etcdLocker) IsMaster(ctx context.Context) (bool, error) {
	resp, err := l.etcd.Get(ctx, l.key)
	if err != nil {
		return false, errors.Wrap(err, "fail to get lock")
	}

	if len(resp.Kvs) != 1 {
		// DAFUK :/
		return false, errors.New("Invalid ETCD state (key not found!)")
	}

	return resp.Kvs[0].Lease == int64(l.leaseID), nil
}

func (l *etcdLocker) Stop(ctx context.Context) error {
	// Reset the lease and let the old lease die.
	// Setting the leaseID to 0 will ensure that the next time `Refresh` is
	// called, we will work with a new lease.
	l.leaseID = 0
	return nil
}