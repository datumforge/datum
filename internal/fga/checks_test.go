package fga

// func Test_CheckDirectUser(t *testing.T) {
// 	ec, err := echox.NewTestContextWithValidUser("funk")
// 	if err != nil {
// 		t.Fatal()
// 	}

// 	echoContext := *ec

// 	ctx := context.WithValue(echoContext.Request().Context(), echox.EchoContextKey, echoContext)

// 	echoContext.SetRequest(echoContext.Request().WithContext(ctx))

// 	url := os.Getenv("TEST_FGA_URL")
// 	if url == "" {
// 		url = defaultFGAURL
// 	}

// 	fc := newTestFGAClient(t, url)

// 	// seed some relations

// 	if err = fc.CreateRelationshipTupleWithUser(ctx, "member", "organization:datum"); err != nil {
// 		t.Fatal()
// 	}

// 	testCases := []struct {
// 		name        string
// 		relation    string
// 		object      string
// 		expectedRes bool
// 		errRes      string
// 	}{
// 		{
// 			name:        "happy path, valid tuple",
// 			relation:    "member",
// 			object:      "organization:datum",
// 			expectedRes: true,
// 			errRes:      "",
// 		},
// 		{
// 			name:        "tuple does not exist",
// 			relation:    "member",
// 			object:      "organization:google",
// 			expectedRes: false,
// 			errRes:      "",
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			valid, err := fc.CheckDirectUser(ctx, tc.relation, tc.object)

// 			if tc.errRes != "" {
// 				assert.Error(t, err)
// 				assert.ErrorContains(t, err, tc.errRes)
// 				assert.Equal(t, tc.expectedRes, valid)

// 				return
// 			}

// 			assert.NoError(t, err)
// 			assert.Equal(t, tc.expectedRes, valid)
// 		})
// 	}
// }
