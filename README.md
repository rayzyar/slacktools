# SlackTools
Tools used to get data from slack

## Installing

### *go get*

    $ go get -u github.com/rayzyar/slacktools
    
## Example
### get slack subteam_id of a list of slack user groups by names
slacktools subteam usergroup1,usergroup2
### get slack subteam_id of a list of slack user groups by names' prefix
slacktools -prefix subteam usergroup-prefix1,usergroup-prefix2