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

### AdminUpdate (admin/account/update) // 1245

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
  "modifiedTime": "2023-04-01T12:00:00Z"
}
```

#### output

```json
  {
  "message": ""
}
```

### Search (/user/search/users) // 611

#### Description

Searches for users.

#### Input

```json
{
  "id": "1 // User ID. ",
  "username": "johndoe // Username. ",
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

### Search (/user/search/posts) // 750

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

### TagQuery (/user/query/tag) // 527

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


### LabelQuery (/user/query/label) // 555

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

### SectionQuery (/user/query/section) // 583

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