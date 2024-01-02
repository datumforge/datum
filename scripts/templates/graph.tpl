extend type Query {
    """
    Look up {{ .Name | ToLower }} by ID
    """
     {{ .Name}}(
        """
        ID of the {{ .Name | ToLower }}
        """
        id: ID!
    ):  {{ .Name }}!
}

extend type Mutation{
    """
    Create a new {{ .Name | ToLower }}
    """
    create{{ .Name }}(
        """
        values of the {{ .Name | ToLower }}
        """
        input: Create{{ .Name }}Input!
    ): {{ .Name }}CreatePayload!
    """
    Update an existing {{ .Name | ToLower }}
    """
    update{{ .Name }}(
        """
        ID of the {{ .Name | ToLower }} 
        """
        id: ID!
        """
        New values for the {{ .Name | ToLower }}
        """
        input: Update{{ .Name }}Input!
    ): {{ .Name }}UpdatePayload!
    """
    Delete an existing {{ .Name | ToLower}}
    """
    delete{{ .Name }}(
        """
        ID of the {{ .Name | ToLower }}
        """
        id: ID!
    ): {{ .Name }}DeletePayload!
}

"""
Return response for create{{ .Name }} mutation
"""
type {{ .Name }}CreatePayload {
    """
    Created {{ .Name | ToLower }}
    """
    {{ .Name }}: {{ .Name }}!
}

"""
Return response for update{{ .Name }} mutation
"""
type {{ .Name }}UpdatePayload {
    """
    Updated {{ .Name | ToLower }}
    """
    {{ .Name }}: {{ .Name }}!
}

"""
Return response for delete{{ .Name }} mutation
"""
type {{ .Name }}DeletePayload {
    """
    Deleted {{ .Name | ToLower }} ID
    """
    deletedID: ID!
}
