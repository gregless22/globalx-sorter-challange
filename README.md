# globalx-sorter-challange

### Testing

Testing can be done by running "go test ./cmd/main.go"
use the -v flag to get verbose output

### Build

To build run "go build -o ./bin/name-sorter cmd/main.go" in the project directory.
This will build an executable file in the bin directory

It is built in Go 1.16.4

### Docker

Build the docker image by running "docker build . -t name-sorter" with any other flags you wish
Unfortunately at the moment it only copies in the "unsorted-names-list.txt" file at build time so if you need to change the list, the image needs to be rebuild. You cannot change the file at runtime.
Run the image with docker run <id> where the id is of the docker image id or use the tag "name-sorter"

### Program function

To use the name-sorter the command is ./bin/name-sorter <filename>
The file should be format as per the example file "unsorted-names-list.txt" in the root project folder.
The name-sorter will sort the list of names, based upon their last name only.
It will then save them to "sorted-names-list.txt" file. WARNING!! this will overwrite previous contents.
