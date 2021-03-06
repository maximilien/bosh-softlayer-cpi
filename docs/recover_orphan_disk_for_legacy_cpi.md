# Recover the missing disk from wrong 'bosh deploy'

# Table of Contents

- [Purpose](#purpose)
- [User Impact](#user_impact)
- [Instructions to Fix](#instructions)
  - [Scenario: user_data.json exists](#scenario1)

# <a id="purpose"></a>Purpose
When updating deployment from losing persistent disk manifest, the old disk will be detached from the updated vm but still alive in orphan disks list(`bosh disks -o`). The missing disk needs to be recovered if we need to keep all data be stored in this disk.

# <a id="user_impact"></a>User Impact
If you don't perform this run book, Orphaned disks are deleted after some days by the director. This will lose all data in this disk.

# <a id="instructions"></a>Instructions to Fix

- ## <a id="scenario"></a>Scenario: user_data.json exists

    Make sure the orphaned disk can still been seen from [SoftLayer Portal](https://control.softlayer.com) and `/var/vcap/bosh/user_data.json` exists on the target vm.

    Take `uaadb/0` as an example of the vm.

    ### 1) Get information and attach the disk to the vm from SoftLayer portal.

    The `vm_cid` can be got from `bosh vms` such as `12345678` and The `disk_cid` can be got from `bosh disks -o` such as `87654321`

    Need to get information from SoftLayer Portal and authorize isntance `12345678`  to disk `87654321`:
    - iSCSI target_address
    - vm's initiator_name
    - vm's username
    - vm's password

    ### 2) Update `/etc/iscsi/initiatorname.iscsi` and `/etc/iscsi/iscsid.conf` in target vm.

    Update `/etc/iscsi/initiatorname.iscsi` and `/etc/iscsi/iscsid.conf` file with these information above and re-discover iSCSI device.
    ```
    # cat > /etc/iscsi/initiatorname.iscsi <<EOF
    InitiatorName={{initiator_name}}
    EOF

    # cat > /etc/iscsi/iscsid.conf <<EOF
    # Generated by bosh-agent
    node.startup = automatic
    node.session.auth.authmethod = CHAP
    node.session.auth.username = {{username}}
    node.session.auth.password = {{password}}
    discovery.sendtargets.auth.authmethod = CHAP
    discovery.sendtargets.auth.username = {{username}}
    discovery.sendtargets.auth.password = {{password}}
    node.session.timeo.replacement_timeout = 120
    node.conn[0].timeo.login_timeout = 15
    node.conn[0].timeo.logout_timeout = 15
    node.conn[0].timeo.noop_out_interval = 10
    node.conn[0].timeo.noop_out_timeout = 15
    node.session.iscsi.InitialR2T = No
    node.session.iscsi.ImmediateData = Yes
    node.session.iscsi.FirstBurstLength = 262144
    node.session.iscsi.MaxBurstLength = 16776192
    node.conn[0].iscsi.MaxRecvDataSegmentLength = 65536
    EOF

    # service open-iscsi restart
    # service multipath-tools restart

    # iscsiadm -m discovery -t sendtargets -p {{target_address}}
    # iscsiadm -m node -l
    ```

    ### 3) Update the user_data.json and restart agent

    Find device path from `dmsetup ls` like `/dev/mapper/3600a098038303845372b4a523234692d` and replace persistent field to `"persistent":{"87654321":"/dev/mapper/3600a098038303845372b4a523234692d"}`

    Restart agent
    ```
    # sv restart agent
    ```

    ### 4) Update "persistent_disks" and "orphan_disks" tables in bosh db
    Login bosh director and update bosh db: delete target disk from `orphan_disks` table and update `persistent_disks` by new record.

    ```
    # cd /var/vcap/packages/postgres/bin
    # ./psql -U postgres -p 5432 bosh

    > select id,job,index from instances order by 2,3;
     id  |         job         | index
    -----+---------------------+-------
      12 | uaadb               |     0

    > select * from orphan_disks;
     id | disk_cid | size  | availability_zone | deployment_name  |                   instance_name                   | cloud_properties_json |         created_at
    ----+----------+-------+-------------------+------------------+---------------------------------------------------+-----------------------+----------------------------
     14 | 87654321 | 20480 | dal12             | cloud-foundry    | uaadb/0                                           | {}                    | 2018-03-12 11:07:19.575882
    > delete from orphan_disks where id =14 ;

    > insert into persistent_disks values (nextval('persistent_disks_id_seq'::regclass),12, 87654321,20480,'t','{}','');
    > select * from persistent_disks order by 1;
     id | instance_id | disk_cid |  size  | active | cloud_properties_json | name
    ----+-------------+----------+--------+--------+-----------------------+------
     50 |          12 | 87654321 |  20480 | t      | {}                    |

    ```

    ### 4) Verify

    Run `bosh cck` and make sure `No problems found`.
    ```
    # bosh cck
    RSA 1024 bit CA certificates are loaded due to old openssl compatibility
    Acting as user 'admin' on deployment 'nc-awhs-us-south' on 'bosh'
    Performing cloud check...
    Director task 108224
      Started scanning 31 vms
      Started scanning 31 vms > Checking VM states. Done (00:00:01)
      Started scanning 31 vms > 31 OK, 0 unresponsive, 0 missing, 0 unbound. Done (00:00:00)
         Done scanning 31 vms (00:00:01)
      Started scanning 15 persistent disks
      Started scanning 15 persistent disks > Looking for inactive disks. Done (00:00:01)
      Started scanning 15 persistent disks > 15 OK, 0 missing, 0 inactive, 0 mount-info mismatch. Done (00:00:00)
         Done scanning 15 persistent disks (00:00:01)
    Task 108224 done
    Started         2018-03-12 15:54:26 UTC
    Finished        2018-03-12 15:54:28 UTC
    Duration        00:00:02
    Scan is complete, checking if any problems found...
    No problems found
    ```

    ### 5) Re-run `bosh deploy` with proper manifest
