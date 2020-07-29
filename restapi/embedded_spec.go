// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API Specification for IMDB service.",
    "title": "API Specification for IMDB service.",
    "contact": {
      "email": "vibhordubey333@gmail.com"
    },
    "license": {
      "name": "MIT",
      "url": "https://github.com/vibhordubey333"
    },
    "version": "1.0.0"
  },
  "host": "localhost:7777",
  "basePath": "/v1",
  "paths": {
    "/imdbservice/addcomment": {
      "post": {
        "description": "Add movie into DB.",
        "tags": [
          "addComment"
        ],
        "summary": "Add movie into DB.",
        "operationId": "postcomments",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "MovieComment": {
                  "description": "Comment for movie.",
                  "type": "string"
                },
                "MovieName": {
                  "description": "Movie name.",
                  "type": "string"
                },
                "UserName": {
                  "description": "Valid username.",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Comment added successfully.",
            "schema": {
              "type": "object",
              "properties": {
                "Code": {
                  "description": "Response Code",
                  "type": "string"
                },
                "Message": {
                  "description": "Comment addded message.",
                  "type": "string"
                }
              }
            }
          },
          "500": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/imdbservice/addmovie": {
      "post": {
        "description": "Add movie into DB.",
        "tags": [
          "addMovie"
        ],
        "summary": "Add movie into DB.",
        "operationId": "postmovie",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "MovieComments": {
                  "description": "Movie comments.",
                  "type": "string"
                },
                "MovieName": {
                  "description": "Movie name.",
                  "type": "string"
                },
                "MovieRating": {
                  "description": "Movie rating.",
                  "type": "string"
                },
                "UserName": {
                  "description": "Provid valid username.",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Movie Uploaded Successfully",
            "schema": {
              "type": "object",
              "properties": {
                "Code": {
                  "description": "Response Code",
                  "type": "string"
                },
                "Message": {
                  "description": "Success message",
                  "type": "string"
                }
              }
            }
          },
          "500": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/imdbservice/addrating": {
      "post": {
        "description": "Add movie into DB.",
        "tags": [
          "addMovie"
        ],
        "summary": "Add movie into DB.",
        "operationId": "addmovierating",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "MovieName": {
                  "description": "Movie name.",
                  "type": "string"
                },
                "MovieRating": {
                  "description": "Movie rating.",
                  "type": "string"
                },
                "UserName": {
                  "description": "Valid username.",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Rating added successfully.",
            "schema": {
              "type": "object",
              "properties": {
                "Code": {
                  "description": "Response Code",
                  "type": "string"
                },
                "Message": {
                  "description": "Message.",
                  "type": "string"
                }
              }
            }
          },
          "500": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/imdbservice/movieinfo": {
      "get": {
        "description": "Get a information about movie.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "movieinfo"
        ],
        "summary": "Get a information about movie.",
        "operationId": "getmoviesinfo",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "MovieName": {
                  "description": "Movie name.",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success.",
            "schema": {
              "type": "object",
              "properties": {
                "AvgRating": {
                  "description": "Rating.",
                  "type": "string"
                },
                "MovieName": {
                  "description": "Response Code",
                  "type": "string"
                },
                "PeopleRated": {
                  "description": "Rating count.",
                  "type": "integer"
                },
                "UserComments": {
                  "type": "array",
                  "items": {
                    "type": "string",
                    "properties": {
                      "comment": {
                        "description": "Comment given by user.",
                        "type": "string"
                      },
                      "username": {
                        "description": "Name of the user.",
                        "type": "string"
                      }
                    }
                  }
                }
              }
            }
          },
          "500": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/imdbservice/moviesrating": {
      "get": {
        "description": "Movies rated by a user.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "ratemovies"
        ],
        "summary": "Movies rated by a user.",
        "operationId": "ratemovies",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "MovieName": {
                  "description": "Movie name.",
                  "type": "string"
                },
                "UserName": {
                  "description": "Movie rating.",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Listing movies insepected by user.",
            "schema": {
              "type": "array",
              "items": {
                "type": "object",
                "properties": {
                  "Comments": {
                    "description": "Comments given on the movie by user.",
                    "type": "string"
                  },
                  "MovieName": {
                    "description": "Name of the movie.",
                    "type": "string"
                  },
                  "Rating": {
                    "description": "Rating of the movie.",
                    "type": "string"
                  }
                }
              }
            }
          },
          "500": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "API Specification for IMDB service.",
      "name": "imdb",
      "externalDocs": {
        "description": "Find out more",
        "url": "https://github.com/vibhordubey333"
      }
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API Specification for IMDB service.",
    "title": "API Specification for IMDB service.",
    "contact": {
      "email": "vibhordubey333@gmail.com"
    },
    "license": {
      "name": "MIT",
      "url": "https://github.com/vibhordubey333"
    },
    "version": "1.0.0"
  },
  "host": "localhost:7777",
  "basePath": "/v1",
  "paths": {
    "/imdbservice/addcomment": {
      "post": {
        "description": "Add movie into DB.",
        "tags": [
          "addComment"
        ],
        "summary": "Add movie into DB.",
        "operationId": "postcomments",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "MovieComment": {
                  "description": "Comment for movie.",
                  "type": "string"
                },
                "MovieName": {
                  "description": "Movie name.",
                  "type": "string"
                },
                "UserName": {
                  "description": "Valid username.",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Comment added successfully.",
            "schema": {
              "type": "object",
              "properties": {
                "Code": {
                  "description": "Response Code",
                  "type": "string"
                },
                "Message": {
                  "description": "Comment addded message.",
                  "type": "string"
                }
              }
            }
          },
          "500": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/imdbservice/addmovie": {
      "post": {
        "description": "Add movie into DB.",
        "tags": [
          "addMovie"
        ],
        "summary": "Add movie into DB.",
        "operationId": "postmovie",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "MovieComments": {
                  "description": "Movie comments.",
                  "type": "string"
                },
                "MovieName": {
                  "description": "Movie name.",
                  "type": "string"
                },
                "MovieRating": {
                  "description": "Movie rating.",
                  "type": "string"
                },
                "UserName": {
                  "description": "Provid valid username.",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Movie Uploaded Successfully",
            "schema": {
              "type": "object",
              "properties": {
                "Code": {
                  "description": "Response Code",
                  "type": "string"
                },
                "Message": {
                  "description": "Success message",
                  "type": "string"
                }
              }
            }
          },
          "500": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/imdbservice/addrating": {
      "post": {
        "description": "Add movie into DB.",
        "tags": [
          "addMovie"
        ],
        "summary": "Add movie into DB.",
        "operationId": "addmovierating",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "MovieName": {
                  "description": "Movie name.",
                  "type": "string"
                },
                "MovieRating": {
                  "description": "Movie rating.",
                  "type": "string"
                },
                "UserName": {
                  "description": "Valid username.",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Rating added successfully.",
            "schema": {
              "type": "object",
              "properties": {
                "Code": {
                  "description": "Response Code",
                  "type": "string"
                },
                "Message": {
                  "description": "Message.",
                  "type": "string"
                }
              }
            }
          },
          "500": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/imdbservice/movieinfo": {
      "get": {
        "description": "Get a information about movie.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "movieinfo"
        ],
        "summary": "Get a information about movie.",
        "operationId": "getmoviesinfo",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "MovieName": {
                  "description": "Movie name.",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success.",
            "schema": {
              "type": "object",
              "properties": {
                "AvgRating": {
                  "description": "Rating.",
                  "type": "string"
                },
                "MovieName": {
                  "description": "Response Code",
                  "type": "string"
                },
                "PeopleRated": {
                  "description": "Rating count.",
                  "type": "integer"
                },
                "UserComments": {
                  "type": "array",
                  "items": {
                    "type": "string",
                    "properties": {
                      "comment": {
                        "description": "Comment given by user.",
                        "type": "string"
                      },
                      "username": {
                        "description": "Name of the user.",
                        "type": "string"
                      }
                    }
                  }
                }
              }
            }
          },
          "500": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/imdbservice/moviesrating": {
      "get": {
        "description": "Movies rated by a user.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "ratemovies"
        ],
        "summary": "Movies rated by a user.",
        "operationId": "ratemovies",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "MovieName": {
                  "description": "Movie name.",
                  "type": "string"
                },
                "UserName": {
                  "description": "Movie rating.",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Listing movies insepected by user.",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/RatemoviesOKBodyItems0"
              }
            }
          },
          "500": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "RatemoviesOKBodyItems0": {
      "type": "object",
      "properties": {
        "Comments": {
          "description": "Comments given on the movie by user.",
          "type": "string"
        },
        "MovieName": {
          "description": "Name of the movie.",
          "type": "string"
        },
        "Rating": {
          "description": "Rating of the movie.",
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "API Specification for IMDB service.",
      "name": "imdb",
      "externalDocs": {
        "description": "Find out more",
        "url": "https://github.com/vibhordubey333"
      }
    }
  ]
}`))
}
