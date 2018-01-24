## eet

### Summary 
This is a slack bot that is meant to gather volunteers for people who would like to get lunch with a group of random people within their organization.

### API
POST to `/join_meeting`
POST to `/leave_meeting`

NOTE: The entire `jsonapi` package is essentially a straight copy of the `github.com/go-kit/kit/transport/http` package, updated to make use of OrderMyGear's jsonapi library's encoding and decoding.
