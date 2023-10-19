extend type Query {
    """
    Look up {{.Env.name}} by ID
    """
     {{.Env.name}}(
        """
        ID of the {{.Env.name}}
        """
        id: ID!
    ):  {{.Env.object}}!
}

extend type Mutation{
    """
    Create a new {{.Env.name}}
    """
    create{{.Env.object}}(
        """
        values of the {{.Env.name}}
        """
        input: Create{{.Env.object}}Input!
    ): {{.Env.object}}CreatePayload!
    """
    Update an existing {{.Env.name}}
    """
    update{{.Env.object}}(
        """
        ID of the {{.Env.name}}
        """
        id: ID!
        """
        New values for the {{.Env.name}}
        """
        input: Update{{.Env.object}}Input!
    ): {{.Env.object}}UpdatePayload!
    """
    Delete an existing {{.Env.name}}
    """
    delete{{.Env.object}}(
        """
        ID of the {{.Env.name}}
        """
        id: ID!
    ): {{.Env.object}}DeletePayload!
}

"""
Return response for create{{.Env.object}} mutation
"""
type {{.Env.object}}CreatePayload {
    """
    Created {{.Env.name}}
    """
    {{.Env.name}}: {{.Env.object}}!
}

"""
Return response for update{{.Env.object}} mutation
"""
type {{.Env.object}}UpdatePayload {
    """
    Updated {{.Env.name}}
    """
    {{.Env.name}}: {{.Env.object}}!
}

"""
Return response for delete{{.Env.object}} mutation
"""
type {{.Env.object}}DeletePayload {
    """
    Deleted {{.Env.name}} ID
    """
    deletedID: ID!
}