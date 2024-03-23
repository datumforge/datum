package model

// Merge attempts to fold a spec into the current spec
// In general, the provided spec takes precedence over the current one
func (o *OpenAPI) Merge(spec OpenAPI) {
	o.Servers = append(spec.Servers, o.Servers...)

	for path, obj := range spec.Paths {
		pathObj, ok := spec.Paths[path]
		if !ok {
			pathObj = obj
		} else {
			if pathObj.Delete == nil {
				pathObj.Delete = obj.Delete
			}

			if pathObj.Get == nil {
				pathObj.Get = obj.Get
			}

			if pathObj.Patch == nil {
				pathObj.Patch = obj.Patch
			}

			if pathObj.Post == nil {
				pathObj.Post = obj.Post
			}

			if pathObj.Put == nil {
				pathObj.Put = obj.Put
			}

			if pathObj.Options == nil {
				pathObj.Options = obj.Options
			}

			if pathObj.Trace == nil {
				pathObj.Trace = obj.Trace
			}

			pathObj.Parameters = append(pathObj.Parameters, obj.Parameters...)
			if pathObj.Description == "" {
				pathObj.Description = obj.Description
			} else {
				pathObj.Description += "\n" + obj.Description
			}

			pathObj.Servers = append(spec.Servers, o.Servers...)
			if pathObj.Summary == "" {
				pathObj.Summary = obj.Summary
			} else {
				pathObj.Summary += "\n" + obj.Summary
			}
		}

		o.Paths[path] = pathObj
	}

	if o.Components == nil {
		o.Components = spec.Components
	} else {
		for k, v := range spec.Components.Parameters {
			if _, ok := o.Components.Parameters[k]; !ok {
				o.Components.Parameters[k] = v
			}
		}

		for k, v := range spec.Components.Callbacks {
			if _, ok := o.Components.Callbacks[k]; !ok {
				o.Components.Callbacks[k] = v
			}
		}

		for k, v := range spec.Components.Headers {
			if _, ok := o.Components.Headers[k]; !ok {
				o.Components.Headers[k] = v
			}
		}

		for k, v := range spec.Components.Links {
			if _, ok := o.Components.Links[k]; !ok {
				o.Components.Links[k] = v
			}
		}

		for k, v := range spec.Components.PathItems {
			if _, ok := o.Components.PathItems[k]; !ok {
				o.Components.PathItems[k] = v
			}
		}

		for k, v := range spec.Components.RequestBodies {
			if _, ok := o.Components.RequestBodies[k]; !ok {
				o.Components.RequestBodies[k] = v
			}
		}

		for k, v := range spec.Components.Responses {
			if _, ok := o.Components.Responses[k]; !ok {
				o.Components.Responses[k] = v
			}
		}

		for k, v := range spec.Components.Schemas {
			if _, ok := o.Components.Schemas[k]; !ok {
				o.Components.Schemas[k] = v
			}
		}

		for k, v := range spec.Components.SecuritySchemes {
			if _, ok := o.Components.SecuritySchemes[k]; !ok {
				o.Components.SecuritySchemes[k] = v
			}
		}
	}

	if o.ExternalDocs == nil {
		o.ExternalDocs = spec.ExternalDocs
	}

	o.Security = append(spec.Security, o.Security...)
	o.Tags = append(spec.Tags, o.Tags...)

	if len(o.Webhooks) == 0 {
		o.Webhooks = spec.Webhooks
	} else {
		for k, v := range spec.Webhooks {
			if _, ok := o.Webhooks[k]; !ok {
				o.Webhooks[k] = v
			}
		}
	}
}

// Merge folds a Request object into the current Request object
// In general, the provided request takes precedence over the current one
func (r *RequestSpec) Merge(other RequestSpec) {
	r.Parameters = append(r.Parameters, other.Parameters...)
	if r.RequestBody == nil {
		r.RequestBody = other.RequestBody
	} else if other.RequestBody != nil {
		if len(other.RequestBody.Description) != 0 {
			r.RequestBody.Description = other.RequestBody.Description
		}

		r.RequestBody.Required = other.RequestBody.Required || r.RequestBody.Required

		if r.RequestBody.Content == nil {
			r.RequestBody.Content = other.RequestBody.Content
		} else {
			for k, v := range other.RequestBody.Content {
				r.RequestBody.Content[k] = v
			}
		}
	}
}

// Merge folds an Operation object into the current Operation object
// In general, the provided operation takes precedence over the current one
func (o *Operation) Merge(other Operation) {
	if other.RequestSpec != nil {
		if o.RequestSpec != nil {
			o.RequestSpec.Merge(*other.RequestSpec)
		} else {
			o.RequestSpec = other.RequestSpec
		}
	}

	if o.Tags == nil || len(o.Tags) == 0 {
		o.Tags = other.Tags
	} else if len(other.Tags) != 0 {
		tagMap := make(map[string]struct{})

		for _, tag := range append(o.Tags, other.Tags...) {
			tagMap[tag] = struct{}{}
		}

		o.Tags = make([]string, len(tagMap))
		i := 0

		for t := range tagMap {
			o.Tags[i] = t
			i++
		}
	}

	if len(other.Summary) != 0 {
		o.Summary = other.Summary
	}

	if len(other.Description) != 0 {
		o.Description = other.Description
	}

	if len(other.OperationID) != 0 {
		o.OperationID = other.OperationID
	}

	o.Deprecated = o.Deprecated || other.Deprecated
	if other.ExternalDocs != nil {
		o.ExternalDocs = other.ExternalDocs
	}

	o.Servers = append(o.Servers, other.Servers...)
	o.Responses.Merge(other.Responses)

	for k, cb := range other.Callbacks {
		if v, ok := o.Callbacks[k]; ok {
			for p, path := range cb {
				v[p] = path
			}
		} else {
			o.Callbacks[k] = cb
		}
	}

	o.Security = append(o.Security, other.Security...)
}

// Merge combines 2 responses objects into a single map
func (r *Responses) Merge(other Responses) {
	for k, v := range other {
		(*r)[k] = v
	}
}
