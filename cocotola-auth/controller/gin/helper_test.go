package controller_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
	"github.com/stretchr/testify/require"

	rsuserdomain "github.com/kujilabo/cocotola-1.23/redstart/user/domain"
)

func organizationID(t *testing.T, organizationID int) *rsuserdomain.OrganizationID {
	t.Helper()
	id, err := rsuserdomain.NewOrganizationID(organizationID)
	require.NoError(t, err)
	return id
}

func appUserID(t *testing.T, appUserID int) *rsuserdomain.AppUserID {
	t.Helper()
	id, err := rsuserdomain.NewAppUserID(appUserID)
	require.NoError(t, err)
	return id
}

func readBytes(t *testing.T, b *bytes.Buffer) []byte {
	t.Helper()
	respBytes, err := io.ReadAll(b)
	require.NoError(t, err)
	return respBytes
}

func parseJSON(t *testing.T, bytes []byte) interface{} {
	t.Helper()
	obj, err := oj.Parse(bytes)
	require.NoError(t, err)
	return obj
}

func parseExpr(t *testing.T, v string) jp.Expr {
	t.Helper()
	expr, err := jp.ParseString(v)
	require.NoError(t, err)
	return expr
}
