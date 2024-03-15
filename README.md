[<picture><source media="(prefers-color-scheme: dark)" srcset="https://steampipe.io/images/steampipe-color-logo-and-wordmark-with-white-bubble.svg"><source media="(prefers-color-scheme: light)" srcset="https://steampipe.io/images/steampipe-color-logo-and-wordmark-with-white-bubble.svg"><img width="67%" alt="Steampipe Logo" src="https://steampipe.io/images/steampipe-color-logo-and-wordmark-with-white-bubble.svg"></picture>](https://steampipe.io)

[![plugins](https://img.shields.io/badge/apis_supported-140-blue)](https://hub.powerpipe.io/plugins) &nbsp; 
[![slack](https://img.shields.io/badge/slack-2297-blue)](https://turbot.com/community/join) &nbsp;
[![maintained by](https://img.shields.io/badge/maintained%20by-Turbot-blue)](https://turbot.com)


[Steampipe](https://steampipe.io) is **the zero-ETL way** to query APIs and services. Use it to expose data sources to SQL.

We offer these Steampipe distributions:

**Steampipe CLI**. Run [queries](https://steampipe.io/docs/query/overview) that translate APIs to tables in the Postgres instance that's bundled with Steampipe.

**Steampipe Postgres FDWs**. Use [native Postgres Foreign Data Wrappers](https://steampipe.io/docs/steampipe_postgres/overview) to translate APIs to foreign tables.

**Steampipe SQLite extensions**. Use [SQLite extensions](https://steampipe.io/docs/steampipe_sqlite/overview) to translate APIS to SQLite virtual tables.

**Steampipe export tools**. Use [standalone binaries](https://steampipe.io/docs/steampipe_export/overview) that export data from APIs, no database required.

**Turbot Pipes**. Use [Turbot Pipes](https://turbot.com/pipes) to run Steampipe in the cloud.

## Demo time!

<img alt="steampipe demo" width=500 src="https://steampipe.io/images/steampipe-sql-demo.gif" >

## Install Steampipe

 The <a href="https://steampipe.io/downloads">downloads</a> page shows you how but tl;dr:
 
Linux or WSL

```sh
sudo /bin/sh -c "$(curl -fsSL https://steampipe.io/install/steampipe.sh)"
```

MacOS

```sh
brew tap turbot/tap
brew install steampipe
```

Now, [install a plugin and run your first query →](https://steampipe.io/docs)

## Steampipe plugins

The Steampipe community has grown a suite of [plugins](https://hub.powerpipe.io/plugins) that map APIs to database tables. Plugins are available for [AWS](https://hub.steampipe.io/plugins/turbot/aws), [Azure](https://hub.steampipe.io/plugins/turbot/azure), [GCP](https://hub.steampipe.io/plugins/turbot/gcp), [Kubernetes](https://hub.steampipe.io/plugins/turbot/kubernetes), [GitHub](https://hub.steampipe.io/plugins/turbot/github), [Microsoft 365](https://hub.steampipe.io/plugins/turbot/microsoft365), [Salesforce](https://hub.steampipe.io/plugins/turbot/salesforce), and many more.

There are more than 2000 tables in all, each clearly documented with copy/paste/run examples.

## Developing

If you want to help develop the core Steampipe binary, these are the steps to build it.

<details>
<summary>Clone</summary>

```sh
git clone git@github.com:turbot/steampipe
```
</details>

<details>
<summary>Build</summary>

```
cd steampipe
make
```

The Steampipe binary lands in `/usr/local/bin/steampipe` directory unless you specify an alternate `OUTPUT_DIR`.
</details>

<details>
<summary>Check the version</summary>

```
$ steampipe --version
steampipe version 0.22.0
```
</details>

<details>
<summary>Install a plugin</summary>

```
$ steampipe plugin install steampipe
```
</details>

<details>
<summary>Run your first query</summary>
 
Try it!

```
steampipe query
> .inspect steampipe
+-----------------------------------+-----------------------------------+
| TABLE                             | DESCRIPTION                       |
+-----------------------------------+-----------------------------------+
| steampipe_registry_plugin         | Steampipe Registry Plugins        |
| steampipe_registry_plugin_version | Steampipe Registry Plugin Version |
+-----------------------------------+-----------------------------------+

> select * from steampipe_registry_plugin;
```
</details>

If you're interested in developing [Steampipe plugins](https://hub.steampipe.io), see our [documentation for plugin developers](https://steampipe.io/docs/develop/overview).

## Turbot Pipes

Bring your team to [Turbot Pipes](https://turbot.com/pipes) to use Steampipe together in the cloud.

## Open source and contributing

This repository is published under the [AGPL 3.0](https://www.gnu.org/licenses/agpl-3.0.html) license. Please see our [code of conduct](https://github.com/turbot/.github/blob/main/CODE_OF_CONDUCT.md). Contributors must sign our [Contributor License Agreement](https://turbot.com/open-source#cla) as part of their first pull request. We look forward to collaborating with you!

[Steampipe](https://steampipe.io) is a product produced from this open source software, exclusively by [Turbot HQ, Inc](https://turbot.com). It is distributed under our commercial terms. Others are allowed to make their own distribution of the software, but cannot use any of the Turbot trademarks, cloud services, etc. You can learn more in our [Open Source FAQ](https://turbot.com/open-source).

## Get involved

**[Join #steampipe on Slack →](https://turbot.com/community/join)**


