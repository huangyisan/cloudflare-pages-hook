# cloudflare-pages-hook

Cloudflare Pages service lacks support for deployment status callback notifications. With the help of 'cloudflare-pages-hook,' it makes such callbacks possible.

## How to use?

**Setting Environment Variables**

setting Cloudflare  environment variables and granting API Token **Read Permissions** for Cloudflare Pages
* CF_ACCOUNT_TOKEN
* CF_ACCOUNT_ID
  
Remember to manage your environment variables securely and avoid sharing them publicly, as they provide access to your Cloudflare resources and services.


**Starting the Service**

```api
./cloudflare-pages-hook -n telegram -t "6369xxxxxx:AAyyyyylK0exxxxxxxxxxxxxxx-ki1234B0" -d "-123456789" -s 2m
```

| args | usage                   | example     |
|------|-------------------------|-------------|
| -n   | notification type       | telegram    |
| -t   | notification tool token | secret:code |
| -d   | notification room id    | -123456789  |
| -s   | wait duration           | 2m          |

**Make a request**

```api
curl "127.0.0.1:8080/deployment?project=${project_name}&commitHash=${commit_hash}&branch=${branch}"
```

| query_args   | usage                         | example                                  |
|--------------|-------------------------------|------------------------------------------|
| project_name | cloudflare pages project name | project-A                                |
| commit_hash  | repo commit hash(full hash)   | 511cb11df4b398c8ec935e19992e996271b388f5 |
| branch       | repo branch                   | main                                     |
