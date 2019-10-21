{{ define "index" }}
{{ template "meta"}}
paths:
  /pet:
    post:
      tags:
      - "pet"
      summary: "Add a new pet to the store"
      description: ""
      operationId: "addPet"
      consumes:
      - "application/json"
      - "application/xml"
      produces:
      - "application/xml"
      - "application/json"
      parameters:
      - in: "body"
        name: "body"
        description: "Pet object that needs to be added to the store"
        required: true
        schema:
          type: "object"
          properties:
              username:
                type: "string"
                example: "admin"
              password:
                type: "string"
                example: "111111"
      responses:
        405:
          description: "Invalid input"
{{ end }}
