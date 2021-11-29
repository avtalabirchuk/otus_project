# otus_project
otus image previewer
## how to
### start project 
`make docker-run`
### check
- url to check
`http://127.0.0.1:8080/fill/1600/300/nas-national-prod.s3.amazonaws.com/styles/scale_3840_2160/public/birds/photo/black-vulture_flickr-2-immature_0.jpg`
- check cache
`ls ${PWD}/cache`
### stop project
`make docker-run`
### Test
`make test`
### Lint
`make lint`
### Clear cache
`make clear_cache`