Colored calc

## check if connection succeeds

```
ssh -v -i ~/.ssh/<private-key-filename> -T git@github.com
```
## start ssh-agent
```
eval `ssh-agent -s`
```

## add your key to ssh-agent
```
ssh-add ~/.ssh/<private-key-filename>
```

[reference](https://gist.github.com/xirixiz/b6b0c6f4917ce17a90e00f9b60566278)

## add an annotated git tag

```
git tag -a v1.0.0 -m "some-message-here"
```

## add a lightweight git tag

```
git tag v1.0.0
```

## add a tag to any commit

```
git tag -a v1.2 <commit-hash>

```

[reference](https://www.atlassian.com/git/tutorials/inspecting-a-repository/git-tag#:~:text=A%20best%20practice%20is%20to,data%20for%20a%20public%20release.)