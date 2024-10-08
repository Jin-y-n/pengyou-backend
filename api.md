# API Document

## Overview

This document provides detailed information about all available API endpoints.

### Response Format

Every response is formatted as JSON:

```json
{
  "code": "int   0 -> fail, 1 -> success",
  "data": "object  real data, the `output` of the API",
  "message": "string"
}
```

## User-related Endpoints

### Verify (/user/account/verify)

#### Description

Verifies a user by their email or phone number.

#### Input

```json
{
  "email": "example@example.com     // Email address. ",
  "phone": "1234567890    // Phone number. ",
  "client_ip": "123.123.123.123     // Client IP address. "
}
```

> email and phone must have one not null
> client_ip should not be null

#### Output

```json
{
  "message": ""
}
```

### Register (/user/account/register)

#### Description

Registers a new user.

#### Input

```json
{
  "username": "newuser    // Username. ",
  "password": "securepassword     // Password. ",
  "email": "example@example.com     // Email address. ",
  "phone": "1234567890    // Phone number. ",
  "login_time": "2023-04-01T12:00:00Z     // Last login time.",
  "client_ip": "123.123.123.123     // Client IP address. ",
  "device_info": "iPhone 12     // Device information. ",
  "captcha": "ABCD1234    // Captcha code. "
}
```

username, password, captcha ara all not null
phone and email must have one not null

The name may consist of letters (both uppercase and lowercase), numbers, underscores, hyphens, and can include spaces,
with a length between 4 and 20 characters.
The password must contain at least one digit and one letter, with a length between 8 and 20 characters.


> output

```json
{
  "message": ""
}
```

### Login (user/login)

#### Description

Logs in an existing user.

#### Input

```json
{
  "username": "existinguser     // Username.",
  "password": "securepassword     // Password.",
  "email": "example@example.com     // Email address.",
  "phone": "1234567890    // Phone number."
}
```

username, email, phone must have one null
password must not be null

#### Output

```json
{
  "id": "1    // User unique identifier.",
  "username": "existinguser     // Username.",
  "password": "hashedpassword     // Encrypted password.",
  "email": "example@example.com     // Email address.",
  "phone": "1234567890    // Phone number.",
  "login_time": "2023-04-01T12:00:00Z     // Last login time.",
  "user_profile": {
    "id": "1    // Unique identifier for the user profile.",
    "user_id": "1     // User ID.",
    "display_name": "John Doe     // Display name.",
    "avatar_id": "https:    //example.com/avatar.jpg    // Avatar URL.",
    "bio": "A passionate developer.     // Biography.",
    "gender": "0    // Gender (0 = Male, 1 = Female, 2 = Other).",
    "birthday": "1990-01-01     // Birthday.",
    "location": "New York, NY     // Location.",
    "occupation": "Software Engineer    // Occupation.",
    "education": "Bachelor of Science in Computer Science     // Education background.",
    "school": "University of Technology     // School.",
    "major": "Computer Science    // Major.",
    "company": "TechCorp    // Company.",
    "position": "Senior Developer     // Position.",
    "website": "https:    //example.com/johndoe     // Website.",
    "created_at": "2023-04-01T12:00:00Z     // Creation timestamp.",
    "tag": [
      {
        "id": "1    // Unique identifier for the tag.",
        "name": "Developer    // Tag name."
      },
      {
        "id": "2    // Unique identifier for the tag.",
        "name": "Tech Enthusiast    // Tag name."
      }
    ]
  }
}
```

### Update (/user/update)

#### Description

Updates a user's profile.

#### Input

```json
{
  "id": "1      // User profile unique identifier. ",
  "user_id": "1       // User ID. ",
  "display_name": "John Doe       // Display name. ",
  "avatar_id": "https://example.com/avatar.jpg      // Avatar URL. ",
  "bio": "A passionate developer.       // Biography. ",
  "gender": "0      // Gender (0 = Male, 1 = Female, 2 = Other). ",
  "birthday": "1990-01-01       // Birthday. ",
  "location": "New York, NY       // Location. ",
  "occupation": "Software Engineer        // Occupation. ",
  "education": "Bachelor of Science in Computer Science       // Education background. ",
  "school": "University of Technology       // School. ",
  "major": "Computer Science        // Major. ",
  "company": "TechCorp        // Company. ",
  "position": "Senior Developer       // Position. ",
  "website": "https://example.com/johndoe       // Website. "
}
```

id, user_id must not be null

#### Output

```json
{
  "message": "Profile update successful."
}
```

### UserUpdateSensitiveInfo (/user/update-sensitive)

#### Description

Updates sensitive user information such as passwords and contact details.

#### Input

id, captcha must not be null,
the json from frontend will be checked,
if the captcha is send via email, the phone can not be changed, similarly, if the captcha is send via phone, the email
can not be changed

```json
{
  "id": "1 // User unique identifier. ",
  "username": "newusername // Username. ",
  "password": "newpassword // New password.",
  "email": "newemail@example.com // Email address. ",
  "phone": "9876543210 // Phone number. ",
  "captcha": "ABCD1234 // Captcha code. "
}
```

#### Output

```json
{
  "message": ""
}
```

### UserFriendAdd (/user/friend/add)

#### Description

Adds a new friend to the user's friend list.

#### Input

all fields must be not null

```json
{
  "user_id": "1 // User ID.",
  "friend_id": "2 // Friend ID. ",
  "request_date": "2023-04-01T12:00:00Z // Request date. "
}
```

#### Output

```json
{
  "message": ""
}
```

### UserFriendDelete (/user/friend/delete)

#### Description

Deletes a friend from the user's friend list.

#### Input

all fields must not be null

```json
{
  "id": "1 // Friendship ID. ",
  "user_id": "1 // User ID.",
  "friend_id": "2 // Friend ID. "
}
```

#### output

```json
{
  "message": ""
}
```

### UserFriendAccept (user/friend/accept)

#### Description

Accepts a friend request.

#### Input

all must be not null

```json
{
  "id": "1 // Friendship ID. ",
  "user_id": "1 // User ID. ",
  "friend_id": "2 // Friend ID. ",
  "status": "1 // Status (0 = Pending, 1 = Accepted, 2 = Rejected). "
}
```

#### output

```json
{
  "message": ""
}
```

### UserFriendUpgrade (user/friend/upgrade)

#### Description

Upgrades a friend's relationship.

#### Input

all must be not null

```json
{
  "id": "1 // Friendship ID. ",
  "user_id": "1 // User ID. ",
  "friend_id": "2 // Friend ID. ",
  "relationship": "1 // Relationship (1 = Normal Friend, 2 = Gossip, 3 = Dead, 4 = Couple). "
}
```

#### Output

```json
{
  "message": ""
}
```

### UserSocialAccountAdd (user/friend/social-account/add)

#### Description

Adds a social account to the user's profile.

#### Input

```json
{
  "user_id": "1 // User ID. ",
  "platform": "platform // Platform name. ",
  "link": "https://example.com/johndoe // Link. "
}
```

all must not be null

#### output

```json
{
  "message": ""
}
```

### UserSocialAccountDelete (user/friend/social-account/delete)

#### Description

Deletes a social account from the user's profile.

#### Input

```json
{
  "id": "1 // Social account ID. ",
  "user_id": "1 // User ID. "
}
```

all must not be null

#### Output

```json
{
  "message": ""
}
```

## Post

### PostAdd (/user/post/add)

#### Description

Adds a new post.

#### Input

```json
{
  "author": "1 // Author ID. ",
  "title": "New Post // Title. ",
  "content": "This is a new post. // Content. ",
  "status": "1 // Status (0 = Draft, 1 = Published). ",
  "labels": [
    {
      "id": "1 // Label ID. ",
      "label": "Developer // Label name"
    }
  ],
  "sections": [
    {
      "id": "1 // Section ID. ",
      "section": "Development // Section name"
    }
  ]
}
```

author, title, content, status, section must be not null
label can be empty but not null

#### Output

```json
{
  "message": ""
}
```

### PostUpdate (/user/post/update)

#### Description

Updates a post.

#### Input

```json
{
  "id": "1 // Post ID. ",
  "author": "1 // Author ID. ",
  "title": "Updated Post // Title. ",
  "content": "This is an updated post. // Content. ",
  "status": "1 // Status (0 = Draft, 1 = Published). ",
  "labels": [
    {
      "id": "1 // Label ID. ",
      "label": "Developer // Label name"
    }
  ],
  "sections": [
    {
      "id": "1 // Section ID. ",
      "section": "Development // Section name"
    }
  ]
}
```

#### Output

```json
{
  "message": ""
}
```

### PostDelete (/user/post/delete)

#### Description

Deletes a post.

#### Input

```json
{
  "ids": [
    0
  ]
}
```

id must not be null

#### Output

```json
{
  "message": ""
}
```

### PostQuery (/user/post/query)

#### Description

Queries posts published by himself/herself.1

#### Input

```json
{
  "id": 0,
  "author": 0,
  "title": "string",
  "content": "string",
  "minCreatedAt": "2023-04-01T12:00:00Z",
  "maxCreatedAt": "2023-04-01T12:00:00Z",
  "minModifiedAt": "2023-04-01T12:00:00Z",
  "maxModifiedAt": "2023-04-01T12:00:00Z",
  "labelsIds": [
    0
  ],
  "labelName": [
    "string"
  ],
  "sectionsIds": [
    0
  ],
  "sectionName": [
    "string"
  ]
}
```

author must not be null, create_time_order_rule default is ASC

#### output

```json
[
  {
    "id": "1 // Post ID. ",
    "author": "1 // Author ID. ",
    "title": "New Post // Title. ",
    "content": "This is a new post. // Content. ",
    "status": "1 // Status (0 = Draft, 1 = Published). ",
    "labels": [
      {
        "id": "1 // Label ID. "
      }
    ],
    "sections": [
      {
        "id": "1 // Section ID. "
      }
    ]
  }
]
```

### TagQuery (/user/query/tag) 

#### Description

Queries tags.

#### Input

```json
{
  "id": "1 // Tag ID. ",
  "name": "Tag // Name. "
}
```

#### Output

```json
[
  {
    "id": "1 // Tag ID. ",
    "name": "name // Tag name",
    "description": "description // Tag description"
  }
]

```

### LabelQuery (/user/query/label)

#### Description

Queries labels.

#### Input

```json
{
  "id": "1 // Label ID. ",
  "label": "Label // Name. "
}
```

#### Output

```json
[
  {
    "id": "1 // Label ID. ",
    "label": "name // Label name",
    "description": "description // Label description"
  }
]

```

### SectionQuery (/user/query/section)

#### Description

Queries sections.

#### Input

```json
{
  "id": "1 // Section ID. ",
  "section": "Section // Name. "
}
```

#### Output

```json
[
  {
    "id": "1 // Section ID. ",
    "section": "name // Section name",
    "description": "description // Section description"
  }
]

```

### Search (/user/search/users)

#### Description

Searches for users.

#### Input

```json
{
  "id": "1 // User ID. ",
  "username": "username // Username. ",
  "email": "johndoe@example.com // Email. ",
  "phone": "1234567890 // Phone number. ",
  "tags": [
    {
      "id": "1 // Tag ID. ",
      "name": "Tag // Name. "
    }
  ]
}
```

#### Output

```json

[
  {
    "id": "1    // User unique identifier.",
    "loginTime": "2023-04-01T12:00:00Z     // Last login time.",
    "user_profile": {
      "id": "1    // Unique identifier for the user profile.",
      "userId": "1     // User ID.",
      "displayName": "John Doe     // Display name.",
      "avatarId": "https:    //example.com/avatar.jpg    // Avatar URL.",
      "bio": "A passionate developer.     // Biography.",
      "gender": "0    // Gender (0 = Male, 1 = Female, 2 = Other).",
      "occupation": "Software Engineer    // Occupation.",
      "education": "Bachelor of Science in Computer Science     // Education background.",
      "school": "University of Technology     // School.",
      "major": "Computer Science    // Major.",
      "website": "https:    //example.com/johndoe     // Website."
    },
    "tags": [
      {
        "id": "1    // Unique identifier for the tag.",
        "name": "Developer    // Tag name."
      }
    ]
  }
]

```

### Search (/user/search/tags)

#### Description

Searches for tags.

#### Input

```json
{
  "id": "1 // Tag ID. ",
  "name": "Tag // Name. ",
  "description": "description // Tag description"
}
```

#### Output

```json
[
  {
    "id": "1 // Tag ID. ",
    "name": "Tag // Name. ",
    "description": "description // Tag description"
  }
]
```

### Search (/user/search/labels)

#### Description

Searches for users, posts, tags, labels, and sections.

#### Input

```json
{
  "id": "1 // Tag ID. ",
  "name": "Tag // Name. ",
  "description": "description // Tag description"
}
```

#### Output

```json
[
  {
    "id": "1 // Tag ID. ",
    "name": "Tag // Name. ",
    "description": "description // Tag description"
  }
]
```

### Search (/user/search/sections)

#### Description

Searches for sections.

#### Input

```json
{
  "id": "1 // Tag ID. ",
  "name": "Tag // Name. ",
  "description": "description // Tag description"
}
```

#### Output

```json
[
  {
    "id": "1 // Tag ID. ",
    "name": "Tag // Name. ",
    "description": "description // Tag description"
  }
]
```

### Search (/user/search/posts)

#### Description

Searches for posts.

#### Input

```json
{
  "id": 0,
  "author": 0,
  "title": "string",
  "content": "string",
  "labelsIds": [
    0
  ],
  "labelName": [
    "string"
  ],
  "sectionsIds": [
    0
  ],
  "sectionName": [
    "string"
  ]
}
```

#### Output

```json
[
  {
    "id": 0,
    "author": 0,
    "title": "string",
    "content": "string",
    "createdAt": "2023-04-01T12:00:00Z",
    "modifiedAt": "2023-04-01T12:00:00Z",
    "labels": {
      "id": 0,
      "label": "string"
    },
    "sections": {
      "id": 0,
      "section": "string"
    }
  }
]
```

### ChatConnectionEstablish (/user/chat/establish)

#### Description

Establishes a chat connection between two users.

#### Input

```json
{
  "user_id": "1 // User ID. ",
  "recipient_id": "2 // Recipient ID. "
}
```

all must not be empty

#### output

web-socket

### ChatMessageSend (via socket)

#### Description

Sends a chat message.

#### Input

```json
{
  "sender_id": "1 // Sender ID. ",
  "recipient_id": "2 // Recipient ID. ",
  "content": "Hello, World! // Message content. ",
  "send_at": "2023-04-01T12:00:00Z // Send time. "
}
```

sender_id, recipient_id, content, send_at must be not empty

#### Output

```json
{
  "message": ""
}
```

### ChatMessageReceive (via socket)

#### receive

listening the websocket and receive the message

```json
{
  "id": "1 // Message ID. ",
  "sender_id": "1 // Sender ID. ",
  "recipient_id": "2 // Recipient ID. ",
  "content": "Hello, World! // Message content. ",
  "sent_at": "2023-04-01T12:00:00Z // Send time. "
}
```

# ADMIN-relative

### AdminRegister (admin/account/register) // 862

#### Description

Add an admin.

#### input

```json
  {
  "username": "string",
  "password": "string",
  "email": "example@exxample.com",
  "phone": "1234567890",
  "createdTime": "2023-04-01T12:00:00Z",
  "modifiedTime": "2023-04-01T12:00:00Z",
  "role": 0,
  "captcha": "6-digit captcha"
}
```

#### output

```json
  {
  "message": ""
}
```

### AdminLogin (admin/account/login)

#### Description

Login as an admin.

#### input

```json
{
  "username": "string",
  "password": "string",
  "email": "example@exxample.com",
  "phone": "1234567890"
}
```

#### output

```json


{
  "user": {
    "id": 0,
   "username": "string",
   "password": "string",
   "email": "example@exxample.com",
   "phone": "1234567890",
   "role": 0
  },
  "jwt": "string"
}
```

### AdminLogout (admin/account/logout)

#### Description

Logout as an admin.

#### input

```json
  {
  "id": 0
}
```

#### output

```json
{
  "message": ""
}
```

### AdminUpdate (admin/account/update)

#### Description

Update an admin.

#### input

```json
{
  "id": 0,
  "username": "string",
  "password": "string",
  "email": "example@example.com",
  "phone": "1234567890"
}
```

#### output

```json
{
  "message": ""
}

```

### AdminDelete (admin/account/delete)

#### Description

Delete an admin.

#### input

```json
  {
  "ids": [
    0
  ]
}
```

#### output

```json
  {
  "message": ""
}
```

### AdminUpdate (admin/account/update)

#### Description


Update an admin.
must be a complete object


#### input

```json
  {
  "id": 0,
  "username": "string",
  "password": "string",
  "email": "example@example.com",
  "phone": "1234567890",
  "modifiedTime": "2023-04-01T12:00:00Z",
  "modifiedPerson": 0,
  "role": 0,
  "modifiedByRoot": 0
}
```

#### output

```json
  {
  "message": ""
}
```

### AdminQuery (admin/account/query)

#### Description

Queries admins.

#### input

```json
  {
  "id": 0,
  "username": "string",
  "password": "string",
  "email": "example@example.com",
  "phone": "1234567890",
  "minCreatedTime": "2023-04-01T12:00:00Z",
  "maxCreatedTime": "2023-04-01T12:00:00Z",
  "minModifiedTime": "2023-04-01T12:00:00Z",
  "maxModifiedTime": "2023-04-01T12:00:00Z",
  "createdPerson": 0,
  "modifiedPerson": 0,
  "role": 0
}
```

#### output

```json
  {
  "id": 0,
  "username": "string",
  "password": "string",
  "email": "example@example.com",
  "phone": "1234567890",
  "createdTime": "2023-04-01T12:00:00Z",
  "modifiedTime": "2023-04-01T12:00:00Z",
  "createdPerson": 0,
  "modifiedPerson": 0,
  "deleteAt": null,
  "role": 0
}
```

### PostSectionAdd (/admin/post-section/add)

#### Description

Add a post section.

#### input

```json
{
  "section": "section // Section name. ",
  "description": "description // Description. "
}
```

#### output

```json
{
  "message": ""
}
```

### PostSectionDelete (/admin/post-section/delete)

#### Description

Delete a post section.

#### input

```json
{
  "ids": [
    0
  ]
}
```

id must not be null

#### output

```json
{
  "message": ""
}
```

### PostSectionQuery (/admin/post-label/section)

#### Description

Add a post label

#### input

```json
{
  "id": 0,
  "section": "section // section. "
}
```

label must not be null

#### output

```json
{
  "id": 0,
  "section": "label // Label. ",
  "description": "description // Description. "
}
```

### PostLabelAdd (/admin/post-label/add)

#### Description

Add a post label

#### input

```json
{
  "label": "label // Label. ",
  "description": "description // Description. "
}
```

label must not be null

#### output

```json
{
  "message": ""
}
```

### PostLabelDelete (/admin/post-label/delete)

#### Description

Delete a post label

#### Input

```json
{
  "ids": [
    0
  ]
}
```

id must not be null

#### output

```json
{
  "message": ""
}
```

### PostLabelQuery (/admin/post-label/query)

#### Description

Add a post label

#### input

```json
{
  "id": 0,
  "label": "label // Label. "
}
```

label must not be null

#### output

```json
{
  "id": 0,
  "label": "label // Label. ",
  "description": "description // Description. "
}
```

### PostQuery (admin/post/query)

#### Description

Queries posts.

#### Input

```json
{
  "id": 0,
  "author": 0,
  "title": "string",
  "content": "string",
  "minCreatedAt": "2023-04-01T12:00:00Z",
  "maxCreatedAt": "2023-04-01T12:00:00Z",
  "minModifiedAt": "2023-04-01T12:00:00Z",
  "maxModifiedAt": "2023-04-01T12:00:00Z",
  "status": 0,
  "createdPerson": 0,
  "modifiedPerson": 0,
  "labelIds": [
    0
  ],
  "labelNames": [
    "string"
  ],
  "sectionIds": [
    0
  ],
  "sectionNames": [
    "string"
  ]
}

```

#### output

```json
[
  {
    "id": 0,
    "author": 0,
    "title": "string",
    "content": "string",
    "createdAt": "2023-04-01T12:00:00Z",
    "modifiedAt": "2023-04-01T12:00:00Z",
    "status": 0,
    "createdPerson": 0,
    "modifiedPerson": 0,
    "labels": [
      {
        "id": 0,
        "label": "string"
      }
    ],
    "sections": [
      {
        "id": 0,
        "section": "string"
      }
    ],
    "deleteAt": null
  }
]
```

### PostUpdate (admin/post/update)

#### Description

update posts.
must a complete post object

#### Input

```json
{
  "id": 0,
  "author": 0,
  "title": "string",
  "content": "string",
  "status": 0,
  "modifiedAt": "2024-08-20T06:04:53Z",
  "modifiedPerson": 0,
  "labels": [
    {
      "id": 0,
      "label": "string"
    }
  ],
  "sections": [
    {
      "id": 0,
      "section": "string"
    }
  ]
}

```

#### output

```json
{
  "message": ""
}
```

### PostDelete (admin/post/delete)

#### Description

Deletes a post.

#### Input

```json
{
  "ids": [
    0
  ]
}
```

#### output

```json
   {
  "message": ""
}













