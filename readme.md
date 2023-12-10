Colored calc

# check if connection succeeds
ssh -v -i <ssh-private-key> -T git@github.com

# start ssh-agent
eval `ssh-agent -s`

# add your key to ssh-agent
ssh-add ~/.ssh/faaltu