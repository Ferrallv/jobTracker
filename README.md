# jobTracker

## Overview
A simple browser-run application for tracking job applications.

The app is run in the browser at url http://localhost:8080/. It is a basic CRUD application to give a GUI for keeping track of applications. 

Here's what it looks like:


Adding an application:
<p align="center">
  <img height="400" width="600" src="https://github.com/Ferrallv/jobTracker/blob/master/readme_images/01_jobtracker_readme.gif">
</p>


Adding a contact:
<p align="center">
  <img height="400" width="600" src="https://github.com/Ferrallv/jobTracker/blob/master/readme_images/02_jobtracker_readme.gif">
</p>


Adding an interview:
<p align="center">
  <img height="400" width="600" src="https://github.com/Ferrallv/jobTracker/blob/master/readme_images/03_jobtracker_readme.gif">
</p>

## Use

After setting up the [Postgresql database](https://ferrallv.github.io/site/project/PostgresForJobtracker/) copy or fork this repo. 

Create a `.env` file and write 

	DATABASE_URL=postgresql://<username>:<password>@localhost:5432/jobtracker

where <username> and <password> are those that you set when creating the database.

Finally to get it started, cd to the directory you copied the repo to and run
	
	go run main.go

This assumes you have [go](https://golang.org/dl/) installed!

---

MIT License

Copyright (c) 2020 Vanin Ferrall

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
