package emails

import (
	sg "github.com/datumforge/datum/internal/utils/sendgrid"
)

// AddContact adds a contact to SendGrid, adding them to the Datum signup marketing list if
// it is configured. This is an upsert operation so existing contacts will be updated.
// The caller can optionally specify additional lists that the contact should be added
// to. If no lists are configured or specified, then the contact is added or updated in
// SendGrid but is not added to any marketing lists - the intent of this is to track within Sendgrid the
// signups and track PLG-related stuff
func (m *EmailManager) AddContact(contact *sg.Contact, listIDs ...string) (err error) {
	if !m.conf.Enabled() {
		return ErrSendgridNotEnabled
	}

	// Setup the request data
	sgdata := &sg.AddContactData{
		Contacts: []*sg.Contact{contact},
	}

	// Add the contact to the specified marketing lists
	if m.conf.DatumListID != "" {
		sgdata.ListIDs = append(sgdata.ListIDs, m.conf.DatumListID)
	}

	for _, listID := range listIDs {
		if listID != "" {
			sgdata.ListIDs = append(sgdata.ListIDs, listID)
		}
	}

	// Invoke the SendGrid API to add the contact
	if err = sg.AddContacts(m.conf.SendGridAPIKey, sgdata); err != nil {
		return err
	}

	return nil
}
