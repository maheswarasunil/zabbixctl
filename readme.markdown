# zabbixctl

**zabbixctl** is tool for working with zabbix server api using command line
interface, it provides effective way for operating on statuses of triggers,
hosts latest data and groups of users.

![dashboard](http://i.imgur.com/0WZkMN0.gif)

## Installation

```
go get github.com/maheswarasunil/zabbixctl
```

afterwards executable will be placed as `$GOPATH/bin/zabbixctl`

## Configuration

**zabbixctl** must be configurated before using, configuration file should be
placed in `~/.config/zabbixctl.conf` and must be written using following
syntax:

```toml
[server]
  address  = "zabbix.local"
  username = "admin"
  password = "password"
  ignoreServerCert = "true" or "false"

[session]
  path = "~/.cache/zabbixctl.session"
```

**zabbixctl** will authorize in 'zabbix.local' server using given user
credentials and save a zabbix session to a file `~/.cache/zabbixctl.session`
and at second run will use saved session instead of new authorization, by the
way zabbix sessions have a ttl that by default equals to 15 minutes, so if
saved zabbix session is outdated, **zabbixctl** will repeat authorization and
rewrite the session file.

## Usage

#####  -T --triggers
Search on zabbix triggers statuses. Triggers could be filtered using
/<pattern> argument, for example, search and acknowledge all triggers in a
problem state and match the word 'cache':
```
  zabbixctl -Tp /cache
```

##### -y --only-nack
Show only not acknowledged triggers.

##### -x --severity
Specify minimum trigger severity.  Once for information, twice for
warning, three for disaster, four for high, five for disaster.

##### -p --problem
Show triggers that have a problem state.

##### -r --recent
Show triggers that have recently been in a problem state.

##### -s --since <date>
Show triggers that have changed their state after the given time, default: 7
days ago

##### -u --until <date>
Show triggers that have changed their state before the given time.

##### -m --maintenance
Show hosts in maintenance.

##### -i --sort <fields>
Show triggers sorted by specified fields, default: lastchange,priority.

##### -o --order <order>
Show triggers in specified order, default: DESC.

##### -n --limit <amount>
Show specified amount of triggers.

##### -k --acknowledge
Acknowledge all retrieved triggers.

##### -f --noconfirm
Do not prompt acknowledge confirmation dialog.

#####  -L --latest-data
Search and show latest data for specified host(s). Hosts can be searched using
wildcard character '*'.  Latest data can be filtered using /<pattern> argument,
for example retrieve latest data for database nodes and search information
about replication:

```
zabbixctl -L dbnode* /replication
```

##### -g --graph
Show links on graph pages.

#####  -G --groups
Search and operate on configuration of users groups.

##### -l --list
Show list users in specified users group.

##### -a --add
Add specified <user> to specified users group.

##### -r --remove
Remove specified <user> from speicifed users group.

##### -f --noconfirm
Do not prompt confirmation dialog.

##### -w --stacked | -b --normal
Returns single link which points to the stacked or normal graph for matched
items.

## Examples

### Listing triggers in a problem state

```
zabbixctlp -Tp
```

### Listing triggers that have recenty been in a problem state

```
zabbixctlp -Tr
```

### Listing and filtering triggers that contain a word mysql

```
zabbixctlp -T /mysql
```

### Listing and acknowledging triggers that severity level is DISASTER

```
zabbixctl -T -xxxxx -k
```

### Listing latest data for db nodes and filtering for information about replication lag

```
zabbixctl -L dbnode* /lag
```

### Opening stacked graph for CPU quote use of selected containers

```
zabbixctl -L 'container-*' /cpu quota --stacked
```

### Listing users groups that starts with 'HTTP_'

```
zabbixctl -G HTTP_*
```

### Listing users groups that contain user admin

```
zabbixctl -G /admin
```

### Adding user admin to groups that contain user guest

```
zabbixctl -G /guest -a admin
```

## License

MIT.
