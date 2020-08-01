# What can I do?
World's issues solutions ranking

## Frontend
- [ ] Static generated nuxt app, I have to rewrite this to use localStorage instead of Cookies
- [ ] Pagespeed 90+
- [ ] CDN

## Main Page
- [x] Header: "What can I do with:"
- [x] Search field
- [ ] Most popular problems rank - Top 10 (After MVP)

## Problem View Page
Header: "Coronavirus" 
- Remove btn only for me
- Ideas ranking
- [x] Show problem's name and description
- [x] Show ideas
- [x] Sort ideas by ratings
- [x] Possible to add an idea
- [x] Possible to rate idea +/-
Idea:
- [x] What can I do?
- [x] Effects
- [ ] Source/s
- [x] Vote up
- [x] Vote down
- [x] Average money price
- [x] Average time price

- [x] MODERATOR: View - Tab with not accepted ideas
- [ ] MODERATOR: Logic - Tab with not accepted ideas
- [ ] MODERATOR: View - Tab with not accepted problems
- [ ] MODERATOR: Logic - Tab with not accepted problems

- [x] Possible to add a problem

## Login page - just login & password
## Dashboard
- [x] Problems waiting for review
- [x] Ideas waiting for review

- [ ] About page with big explaination
- [ ] Take care of SEO Stuff
- [ ] Take care of Lighthouse audits

## Backend
- [x] Safety condition for token without "Bearer" phrase
- [ ] Varnish cache in front of API
- [ ] Autoinvalidating by tags in Varnish
- [x] Vote actions support
- [ ] Moderator's actions support
- [x] Disable cors on dev https://blog.bitsrc.io/how-and-why-you-should-avoid-cors-in-single-page-apps-db25452ad2f8
- [ ] Disable cors on prod https://blog.bitsrc.io/how-and-why-you-should-avoid-cors-in-single-page-apps-db25452ad2f8
- [ ] Load balance - https://codeburst.io/load-balancing-go-api-with-docker-nginx-digital-ocean-d7f05f7c9b31

## Analytic
- [ ] Attach simple google analytics to know what happens there

# Deploy
- [ ] FINALLY!

- [ ] Prepare content
- [ ] 10 Problems
- [ ] 3 x 10 Ideas

## Promotion
- [ ] Medium article
- [ ] Vue discord channel
- [ ] Facebook programmist groups
- [ ] #LECIMY group
- [ ] Linkedin post
- [ ] Facebook feed graphic template for scrollers!!

## Safety before further developing
- [ ] Unit tests for Backend
- [ ] E2E for Frontend
- [ ] Unit tests for Frontend

## Helpers
Import database from file to container:
```
cat backup.sql | docker exec -i <container_id> /usr/bin/mysql -u root --password=qwerty whatcanido
```

As `container_id` put mariadb container ID. You can get it by using `docker ps` command with enabled db.

## How to launch
```
# Launching DB + PHPMyAdmin
 cd server;
 docker-compose up;

# Launching Rest API
 ./air

# Launching PWA
 cd client;
 yarn dev;
```
