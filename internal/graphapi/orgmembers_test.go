package graphapi_test

// func TestQuery_OrgMembers(t *testing.T) {
// 	// setup entdb with authz
// 	authClient := setupAuthEntDB(t)
// 	defer authClient.entDB.Close()

// 	// setup user context
// 	reqCtx, err := userContext()
// 	require.NoError(t, err)

// 	org1 := (&OrganizationBuilder{}).MustNew(reqCtx)

// 	listObjects := []string{fmt.Sprintf("%s:%s", organization, org1.ID)}

// 	testCases := []struct {
// 		name     string
// 		queryID  string
// 		expected *ent.Organization
// 		errorMsg string
// 	}{
// 		{
// 			name:     "happy path, get organization",
// 			queryID:  org1.ID,
// 			expected: org1,
// 		},
// 		{
// 			name:     "invalid-id",
// 			queryID:  "tacos-for-dinner",
// 			errorMsg: "organization not found",
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run("Get "+tc.name, func(t *testing.T) {
// 			mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, true)
// 			mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)

// 			// second check won't happen if org does not exist
// 			if tc.errorMsg == "" {
// 				// we need to check list objects even on a get
// 				// because a parent could be request and that access must always be
// 				// checked before being returned
// 				mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
// 				mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
// 			}

// 			resp, err := authClient.gc.GetOrganizationByID(reqCtx, tc.queryID)

// 			if tc.errorMsg != "" {
// 				require.Error(t, err)
// 				assert.ErrorContains(t, err, tc.errorMsg)
// 				assert.Nil(t, resp)

// 				return
// 			}

// 			require.NoError(t, err)
// 			require.NotNil(t, resp)
// 			require.NotNil(t, resp.Organization)
// 		})
// 	}

// 	// delete created org
// 	(&OrganizationCleanup{OrgID: org1.ID}).MustDelete(reqCtx)
// }
