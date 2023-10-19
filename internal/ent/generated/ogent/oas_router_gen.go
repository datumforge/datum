// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

func (s *Server) cutPrefix(path string) (string, bool) {
	prefix := s.cfg.Prefix
	if prefix == "" {
		return path, true
	}
	if !strings.HasPrefix(path, prefix) {
		// Prefix doesn't match.
		return "", false
	}
	// Cut prefix from the path.
	return strings.TrimPrefix(path, prefix), true
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}

	elem, ok := s.cutPrefix(elem)
	if !ok || len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [1]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'i': // Prefix: "integrations"
				if l := len("integrations"); len(elem) >= l && elem[0:l] == "integrations" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleListIntegrationRequest([0]string{}, elemIsEscaped, w, r)
					case "POST":
						s.handleCreateIntegrationRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET,POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch r.Method {
						case "DELETE":
							s.handleDeleteIntegrationRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "GET":
							s.handleReadIntegrationRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "PATCH":
							s.handleUpdateIntegrationRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "DELETE,GET,PATCH")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/organization"
						if l := len("/organization"); len(elem) >= l && elem[0:l] == "/organization" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleReadIntegrationOrganizationRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					}
				}
			case 'm': // Prefix: "memberships"
				if l := len("memberships"); len(elem) >= l && elem[0:l] == "memberships" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleListMembershipRequest([0]string{}, elemIsEscaped, w, r)
					case "POST":
						s.handleCreateMembershipRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET,POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch r.Method {
						case "DELETE":
							s.handleDeleteMembershipRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "GET":
							s.handleReadMembershipRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "PATCH":
							s.handleUpdateMembershipRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "DELETE,GET,PATCH")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'o': // Prefix: "organization"
							if l := len("organization"); len(elem) >= l && elem[0:l] == "organization" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleReadMembershipOrganizationRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						case 'u': // Prefix: "user"
							if l := len("user"); len(elem) >= l && elem[0:l] == "user" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleReadMembershipUserRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						}
					}
				}
			case 'o': // Prefix: "organizations"
				if l := len("organizations"); len(elem) >= l && elem[0:l] == "organizations" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleListOrganizationRequest([0]string{}, elemIsEscaped, w, r)
					case "POST":
						s.handleCreateOrganizationRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET,POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch r.Method {
						case "DELETE":
							s.handleDeleteOrganizationRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "GET":
							s.handleReadOrganizationRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "PATCH":
							s.handleUpdateOrganizationRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "DELETE,GET,PATCH")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'i': // Prefix: "integrations"
							if l := len("integrations"); len(elem) >= l && elem[0:l] == "integrations" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleListOrganizationIntegrationsRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						case 'm': // Prefix: "memberships"
							if l := len("memberships"); len(elem) >= l && elem[0:l] == "memberships" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleListOrganizationMembershipsRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						}
					}
				}
			case 'u': // Prefix: "users"
				if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleListUserRequest([0]string{}, elemIsEscaped, w, r)
					case "POST":
						s.handleCreateUserRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET,POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch r.Method {
						case "DELETE":
							s.handleDeleteUserRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "GET":
							s.handleReadUserRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "PATCH":
							s.handleUpdateUserRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "DELETE,GET,PATCH")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/memberships"
						if l := len("/memberships"); len(elem) >= l && elem[0:l] == "/memberships" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleListUserMembershipsRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					}
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	summary     string
	operationID string
	pathPattern string
	count       int
	args        [1]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// Summary returns OpenAPI summary.
func (r Route) Summary() string {
	return r.summary
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	elem, ok := s.cutPrefix(elem)
	if !ok {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'i': // Prefix: "integrations"
				if l := len("integrations"); len(elem) >= l && elem[0:l] == "integrations" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "ListIntegration"
						r.summary = "List Integrations"
						r.operationID = "listIntegration"
						r.pathPattern = "/integrations"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						r.name = "CreateIntegration"
						r.summary = "Create a new Integration"
						r.operationID = "createIntegration"
						r.pathPattern = "/integrations"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch method {
						case "DELETE":
							r.name = "DeleteIntegration"
							r.summary = "Deletes a Integration by ID"
							r.operationID = "deleteIntegration"
							r.pathPattern = "/integrations/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "GET":
							r.name = "ReadIntegration"
							r.summary = "Find a Integration by ID"
							r.operationID = "readIntegration"
							r.pathPattern = "/integrations/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "PATCH":
							r.name = "UpdateIntegration"
							r.summary = "Updates a Integration"
							r.operationID = "updateIntegration"
							r.pathPattern = "/integrations/{id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/organization"
						if l := len("/organization"); len(elem) >= l && elem[0:l] == "/organization" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: ReadIntegrationOrganization
								r.name = "ReadIntegrationOrganization"
								r.summary = "Find the attached Organization"
								r.operationID = "readIntegrationOrganization"
								r.pathPattern = "/integrations/{id}/organization"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					}
				}
			case 'm': // Prefix: "memberships"
				if l := len("memberships"); len(elem) >= l && elem[0:l] == "memberships" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "ListMembership"
						r.summary = "List Memberships"
						r.operationID = "listMembership"
						r.pathPattern = "/memberships"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						r.name = "CreateMembership"
						r.summary = "Create a new Membership"
						r.operationID = "createMembership"
						r.pathPattern = "/memberships"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch method {
						case "DELETE":
							r.name = "DeleteMembership"
							r.summary = "Deletes a Membership by ID"
							r.operationID = "deleteMembership"
							r.pathPattern = "/memberships/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "GET":
							r.name = "ReadMembership"
							r.summary = "Find a Membership by ID"
							r.operationID = "readMembership"
							r.pathPattern = "/memberships/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "PATCH":
							r.name = "UpdateMembership"
							r.summary = "Updates a Membership"
							r.operationID = "updateMembership"
							r.pathPattern = "/memberships/{id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'o': // Prefix: "organization"
							if l := len("organization"); len(elem) >= l && elem[0:l] == "organization" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: ReadMembershipOrganization
									r.name = "ReadMembershipOrganization"
									r.summary = "Find the attached Organization"
									r.operationID = "readMembershipOrganization"
									r.pathPattern = "/memberships/{id}/organization"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
						case 'u': // Prefix: "user"
							if l := len("user"); len(elem) >= l && elem[0:l] == "user" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: ReadMembershipUser
									r.name = "ReadMembershipUser"
									r.summary = "Find the attached User"
									r.operationID = "readMembershipUser"
									r.pathPattern = "/memberships/{id}/user"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
						}
					}
				}
			case 'o': // Prefix: "organizations"
				if l := len("organizations"); len(elem) >= l && elem[0:l] == "organizations" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "ListOrganization"
						r.summary = "List Organizations"
						r.operationID = "listOrganization"
						r.pathPattern = "/organizations"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						r.name = "CreateOrganization"
						r.summary = "Create a new Organization"
						r.operationID = "createOrganization"
						r.pathPattern = "/organizations"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch method {
						case "DELETE":
							r.name = "DeleteOrganization"
							r.summary = "Deletes a Organization by ID"
							r.operationID = "deleteOrganization"
							r.pathPattern = "/organizations/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "GET":
							r.name = "ReadOrganization"
							r.summary = "Find a Organization by ID"
							r.operationID = "readOrganization"
							r.pathPattern = "/organizations/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "PATCH":
							r.name = "UpdateOrganization"
							r.summary = "Updates a Organization"
							r.operationID = "updateOrganization"
							r.pathPattern = "/organizations/{id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'i': // Prefix: "integrations"
							if l := len("integrations"); len(elem) >= l && elem[0:l] == "integrations" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: ListOrganizationIntegrations
									r.name = "ListOrganizationIntegrations"
									r.summary = "List attached Integrations"
									r.operationID = "listOrganizationIntegrations"
									r.pathPattern = "/organizations/{id}/integrations"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
						case 'm': // Prefix: "memberships"
							if l := len("memberships"); len(elem) >= l && elem[0:l] == "memberships" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: ListOrganizationMemberships
									r.name = "ListOrganizationMemberships"
									r.summary = "List attached Memberships"
									r.operationID = "listOrganizationMemberships"
									r.pathPattern = "/organizations/{id}/memberships"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
						}
					}
				}
			case 'u': // Prefix: "users"
				if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "ListUser"
						r.summary = "List Users"
						r.operationID = "listUser"
						r.pathPattern = "/users"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						r.name = "CreateUser"
						r.summary = "Create a new User"
						r.operationID = "createUser"
						r.pathPattern = "/users"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch method {
						case "DELETE":
							r.name = "DeleteUser"
							r.summary = "Deletes a User by ID"
							r.operationID = "deleteUser"
							r.pathPattern = "/users/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "GET":
							r.name = "ReadUser"
							r.summary = "Find a User by ID"
							r.operationID = "readUser"
							r.pathPattern = "/users/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "PATCH":
							r.name = "UpdateUser"
							r.summary = "Updates a User"
							r.operationID = "updateUser"
							r.pathPattern = "/users/{id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/memberships"
						if l := len("/memberships"); len(elem) >= l && elem[0:l] == "/memberships" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: ListUserMemberships
								r.name = "ListUserMemberships"
								r.summary = "List attached Memberships"
								r.operationID = "listUserMemberships"
								r.pathPattern = "/users/{id}/memberships"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					}
				}
			}
		}
	}
	return r, false
}
