# Reshow

Reshow cleans file names in the current directory that refer to a TV show (with season and episode).

The current directory is supposed to represent the show title.  
For example, if you are in the `The Blacklist` directory, it renames the `theblacklist-s01e22-final-fr-uselessinfo.avi` file into `The Blacklist S01E22 FINAL FR.avi`.

## Install

```Go
go get -u github.com/arthurwhite/reshow
```

## Usage

1. Go to the show directory.

   ```Shell
   cd ~/Movies/The\ Blacklist
   ```

2. Run Reshow.

   ```Shell
   reshow
   ```
