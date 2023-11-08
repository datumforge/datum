// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"github.com/datumforge/datum/internal/ent/schema","Package":"github.com/datumforge/datum/internal/ent/generated","Schemas":[{"name":"Group","config":{"Table":""},"edges":[{"name":"setting","type":"GroupSettings","unique":true,"required":true},{"name":"users","type":"User"},{"name":"owner","type":"Organization","ref_name":"groups","unique":true,"inverse":true}],"fields":[{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"updated_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0}},{"name":"created_by","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":2,"MixedIn":true,"MixinIndex":0}},{"name":"updated_by","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":3,"MixedIn":true,"MixinIndex":0}},{"name":"id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":2}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"name"}}},{"name":"description","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":"","default_kind":24,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"Skip":8}}},{"name":"logo_url","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"validators":1,"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"Skip":8}}}],"indexes":[{"unique":true,"edges":["owner"],"fields":["name"]}],"hooks":[{"Index":0,"MixedIn":true,"MixinIndex":0}],"policy":[{"Index":0,"MixedIn":true,"MixinIndex":1}],"annotations":{"EntGQL":{"MutationInputs":[{"IsCreate":true},{}],"QueryField":{},"RelayConnection":true}}},{"name":"GroupSettings","config":{"Table":""},"edges":[{"name":"group","type":"Group","ref_name":"setting","unique":true,"inverse":true,"annotations":{"EntGQL":{"Skip":63}}}],"fields":[{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"updated_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0}},{"name":"created_by","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":2,"MixedIn":true,"MixinIndex":0}},{"name":"updated_by","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":3,"MixedIn":true,"MixinIndex":0}},{"name":"id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":1}},{"name":"visibility","type":{"Type":6,"Ident":"groupsettings.Visibility","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"enums":[{"N":"public","V":"PUBLIC"},{"N":"private","V":"PRIVATE"}],"default":true,"default_value":"PUBLIC","default_kind":24,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"join_policy","type":{"Type":6,"Ident":"groupsettings.JoinPolicy","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"enums":[{"N":"open","V":"OPEN"},{"N":"invite_only","V":"INVITE_ONLY"},{"N":"application_only","V":"APPLICATION_ONLY"},{"N":"invite_or_application","V":"INVITE_OR_APPLICATION"}],"default":true,"default_value":"OPEN","default_kind":24,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}}],"hooks":[{"Index":0,"MixedIn":true,"MixinIndex":0}],"annotations":{"EntGQL":{"MutationInputs":[{"IsCreate":true},{}],"QueryField":{},"RelayConnection":true}}},{"name":"Integration","config":{"Table":""},"edges":[{"name":"owner","type":"Organization","ref_name":"integrations","unique":true,"inverse":true}],"fields":[{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"updated_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0}},{"name":"created_by","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":2,"MixedIn":true,"MixinIndex":0}},{"name":"updated_by","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":3,"MixedIn":true,"MixinIndex":0}},{"name":"id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":1}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"name"}}},{"name":"kind","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"immutable":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"kind"}}},{"name":"description","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"secret_name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"immutable":true,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}}],"hooks":[{"Index":0,"MixedIn":true,"MixinIndex":0}],"annotations":{"EntGQL":{"MutationInputs":[{"IsCreate":true},{}],"QueryField":{},"RelayConnection":true}}},{"name":"Organization","config":{"Table":""},"edges":[{"name":"parent","type":"Organization","field":"parent_organization_id","ref":{"name":"children","type":"Organization","annotations":{"EntGQL":{"RelayConnection":true,"Skip":48}}},"unique":true,"inverse":true,"immutable":true},{"name":"users","type":"User","ref_name":"organizations","inverse":true},{"name":"groups","type":"Group","annotations":{"EntSQL":{"on_delete":"CASCADE"}}},{"name":"integrations","type":"Integration","annotations":{"EntSQL":{"on_delete":"CASCADE"}}}],"fields":[{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"updated_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0}},{"name":"created_by","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":2,"MixedIn":true,"MixinIndex":0}},{"name":"updated_by","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":3,"MixedIn":true,"MixinIndex":0}},{"name":"id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":1}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":160,"unique":true,"validators":2,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"name","Skip":8}}},{"name":"description","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"Skip":8}},"comment":"An optional description of the Organization"},{"name":"parent_organization_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"immutable":true,"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"Skip":33,"Type":"ID"},"EntOAS":{"Create":{"Groups":null,"Policy":0},"Delete":{"Groups":null,"Policy":0},"Example":null,"Groups":null,"List":{"Groups":null,"Policy":0},"Read":{"Groups":null,"Policy":0},"ReadOnly":false,"Schema":{"type":"string"},"Skip":false,"Update":{"Groups":null,"Policy":0}}},"comment":"The ID of the parent organization for the organization."}],"hooks":[{"Index":0,"MixedIn":true,"MixinIndex":0}],"annotations":{"EntGQL":{"MutationInputs":[{"IsCreate":true},{}],"QueryField":{},"RelayConnection":true}}},{"name":"RefreshToken","config":{"Table":""},"fields":[{"name":"id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"client_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"nonce","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":2147483647,"validators":1,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"claims_user_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":2147483647,"validators":1,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"claims_username","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":2147483647,"validators":1,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}},{"name":"claims_email","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":2147483647,"validators":1,"position":{"Index":4,"MixedIn":false,"MixinIndex":0}},{"name":"claims_email_verified","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":5,"MixedIn":false,"MixinIndex":0}},{"name":"claims_preferred_username","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":2147483647,"position":{"Index":6,"MixedIn":false,"MixinIndex":0}},{"name":"connector_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":2147483647,"validators":1,"position":{"Index":7,"MixedIn":false,"MixinIndex":0}},{"name":"token","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":2147483647,"position":{"Index":8,"MixedIn":false,"MixinIndex":0}},{"name":"obsolete_token","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":2147483647,"position":{"Index":9,"MixedIn":false,"MixinIndex":0}},{"name":"last_used","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":10,"MixedIn":false,"MixinIndex":0}}]},{"name":"Session","config":{"Table":""},"edges":[{"name":"users","type":"User","unique":true,"comment":"Sessions belong to users"}],"fields":[{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"updated_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0}},{"name":"created_by","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":2,"MixedIn":true,"MixinIndex":0}},{"name":"updated_by","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":3,"MixedIn":true,"MixinIndex":0}},{"name":"id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":1}},{"name":"type","type":{"Type":6,"Ident":"session.Type","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"enums":[{"N":"local","V":"local"},{"N":"oauth","V":"oauth"},{"N":"app_password","V":"app_password"}],"immutable":true,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"comment":"Sessions can derrive from the local (password auth), oauth, or app_password"},{"name":"disabled","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"comment":"The session may be disabled by the user or by automatic security policy"},{"name":"token","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"unique":true,"default":true,"default_kind":19,"immutable":true,"validators":1,"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"comment":"random 32 bytes encoded as base64"},{"name":"user_agent","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":3,"MixedIn":false,"MixinIndex":0},"comment":"The last known user-agent"},{"name":"ips","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":4,"MixedIn":false,"MixinIndex":0},"comment":"All IPs that have been associated with this session. Reverse-chronological order. The current IP is the first item in the slice"}],"indexes":[{"unique":true,"fields":["id"]}],"hooks":[{"Index":0,"MixedIn":true,"MixinIndex":0}],"annotations":{"EntGQL":{"MutationInputs":[{"IsCreate":true},{}],"QueryField":{},"RelayConnection":true}}},{"name":"User","config":{"Table":""},"edges":[{"name":"organizations","type":"Organization"},{"name":"sessions","type":"Session","annotations":{"EntSQL":{"on_delete":"CASCADE"}}},{"name":"groups","type":"Group","ref_name":"users","inverse":true}],"fields":[{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"updated_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0}},{"name":"created_by","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":2,"MixedIn":true,"MixinIndex":0}},{"name":"updated_by","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":3,"MixedIn":true,"MixinIndex":0}},{"name":"id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":2}},{"name":"email","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"unique":true,"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"first_name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":64,"validators":2,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"first_name"}}},{"name":"last_name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":64,"validators":2,"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"last_name"}}},{"name":"display_name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":64,"default":true,"default_value":"unknown","default_kind":24,"validators":2,"position":{"Index":3,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"display_name"}},"comment":"The user's displayed 'friendly' name"},{"name":"locked","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":false,"default_kind":1,"position":{"Index":4,"MixedIn":false,"MixinIndex":0},"comment":"user account is locked if unconfirmed or explicitly locked"},{"name":"avatar_remote_url","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":255,"nillable":true,"optional":true,"validators":2,"position":{"Index":5,"MixedIn":false,"MixinIndex":0},"comment":"URL of the user's remote avatar"},{"name":"avatar_local_file","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":255,"nillable":true,"optional":true,"validators":1,"position":{"Index":6,"MixedIn":false,"MixinIndex":0},"comment":"The user's local avatar file"},{"name":"avatar_updated_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":7,"MixedIn":false,"MixinIndex":0},"comment":"The time the user's (local) avatar was last updated"},{"name":"silenced_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":8,"MixedIn":false,"MixinIndex":0},"comment":"The time the user was silenced"},{"name":"suspended_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":9,"MixedIn":false,"MixinIndex":0},"comment":"The time the user was suspended"},{"name":"recovery_code","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":10,"MixedIn":false,"MixinIndex":0},"sensitive":true,"comment":"local Actor password recovery code generated during account creation"}],"indexes":[{"unique":true,"fields":["id"]}],"hooks":[{"Index":0,"MixedIn":true,"MixinIndex":0}],"policy":[{"Index":0,"MixedIn":true,"MixinIndex":1}],"annotations":{"EntGQL":{"MutationInputs":[{"IsCreate":true},{}],"QueryField":{},"RelayConnection":true}}}],"Features":["sql/versioned-migration","privacy","schema/snapshot","entql","namedges","sql/schemaconfig","intercept","namedges"]}`
