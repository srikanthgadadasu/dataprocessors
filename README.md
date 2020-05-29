# Nadi Data

Future proof data processing

## Copyright headers:

New files that you contribute should use the standard copyright header:

// Copyright 2020 The Nadi Data Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

### Step 1: Clone the Nadi repo

You need to have a local copy of the source checked out from the correct repository. You can check out the Nadi repo onto your local file system anywhere you want

$ git clone https://github.com/srikanthgadadasu/nadi.git

### Step 2: Prepare changes in a new branch

Each Nadi change must be made in a separate branch, created from the master branch. You can use the normal git commands to create a branch and add changes to the staging area:

$ git checkout -b mybranch
$ [edit files...]
$ git add [files...]

### Step 3: Test your changes

You've written and tested your code, but before sending code out for review, run all the tests for the whole tree to make sure the changes don't break other packages or programs

### Step 4: Send changes for review

Once the change is ready and tested over the whole tree, send it for review

#### First line
The first line of the change description is conventionally a short one-line summary of the change, prefixed by the primary affected package.

A rule of thumb is that it should be written so to complete the sentence "This change modifies Go to _____." That means it does not start with a capital letter, is not a complete sentence, and actually summarizes the result of the change.

Follow the first line by a blank line.

#### Main content
The rest of the description elaborates and should provide context for the change and explain what it does. Write in complete sentences with correct punctuation, just like for your comments in Go. Don't use HTML, Markdown, or any other markup language.
