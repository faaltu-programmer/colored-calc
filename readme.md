Colored calc

# check if connection succeeds

```
ssh -v -i ~/.ssh/<private-key-filename> -T git@github.com
```
# start ssh-agent
```
eval `ssh-agent -s`
```

# add your key to ssh-agent
```
ssh-add ~/.ssh/<private-key-filename>
```

[reference](https://gist.github.com/xirixiz/b6b0c6f4917ce17a90e00f9b60566278)