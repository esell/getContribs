# Get Contribs

An overly complex way of giving credit to repo contributors in your README file.
This "app" will query Github for a list of contributors, generate a markdown table with 
thats info and then shove it into your README file wherever you have the string `CONTRIBPOPULATE`. The idea is that this can be used in
your build/deployment pipeline.
Cool right?

```
Usage of ./getContribs:
  -b    Backup Readme
  -f string
        Readme filename (default "README.md")
  -r string
        Github Repo Name
  -u string
        Github Username of Repo Owner
```

