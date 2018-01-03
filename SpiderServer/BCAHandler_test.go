package SpiderServer

import (
	"fmt"
	"testing"
	"time"
)

func TestLoadBCA(t *testing.T) {
	bow, err := BCALogin(UNAME, PWD)
	CheckTestErr(err, t)
	fmt.Print(bow.Title())
	t.Log(bow.Title())

	ParseBCAAccount(bow)
	time.Sleep(time.Second)
	ParseBCATransHis(bow)

	if bow.Url().String() != TOKO_SUCCESS {
		t.Errorf("invalid res[%v]", bow.Url())
	}
}
func TestParseExtInfo(t *testing.T) {
	var testStr = "|\n      \n         \302\240\n            \n         \n      \n      \n      \n   ||\n      \n         \302\240ACCOUNT INFORMATION - BALANCE INQUIRY\n      \n   |\n      |\n   |\n      \n         \n         \n         Account No.\n         \n         \n      \n      \n         \n         \n         Account Type\n         \n         \n      \n      \n         \n         \n         Currency\n         \n      \n      \n         \n         \n         Available Balance\n         \n         \n      \n   |\n      \n         \n         \n         0066941994\n         \n         \n      \n      \n         \n         \n         Tabungan\n         \n         \n      \n      \n         \n         \n         IDR\n         \n      \n      \n         \n         \n         0.00\n         \n         \n      \n   |\n  \n    \n        \n    \n  \n|"
	res := ParseExtInfo(testStr)
	t.Errorf("invalid res[%v]", res)
}

func TestParseTrans(t *testing.T) {
	res := ParseBcaTransBody(transBody)
	for idx, table := range res {
		for idy, item := range table {
			P("%v %v get [%v]\n", idx, idy, item)
		}
	}
	t.Errorf("invalid res[%v]", res)
}

var transBody = `
<layer id="DateTime" left="25" top="12"></layer>
<table border="0" cellpadding="0" cellspacing="0" width="590">
   <tbody><tr height="20" bgcolor="#e7d300">
      <td bgcolor="#e7d300" width="393">
         <font face="Verdana" size="2" color="#0000bb"> 
         </font>
      </td>
      <td width="25" bgcolor="#4a55b5"><img src="https://ibank.klikbca.com/images/latar1b.jpg;bca7548e41af5f81936"/></td>
      <td width="100" bgcolor="#4a55b5" align="center" class="aktif"></td>
   </tr>
   <tr bgcolor="#4A55B5" height="2"><td colspan="3"></td></tr>
</tbody></table>

<table border="0" cellpadding="1" cellspacing="0" width="590">
   <tbody><tr height="20">
      <td colspan="3" bgcolor="#8486de">
         <font size="2" face="Arial,Helvetica,Geneva,Swiss,SunSans-Regular" color="white"> ACCOUNT INFORMATION - ACCOUNT STATEMENT</font>
      </td>
   </tr>
   <tr>
      <td bgcolor="#FFFFFF" colspan="3"></td></tr><tr><td bgcolor="#FFFFFF" colspan="3"></td>
   </tr>
</tbody></table>

<table border="0" cellpadding="0" cellspacing="0" width="590">
<tbody><tr>  <td colspan="2" align="center">    <table border="0" width="90%" cellpadding="0" cellspacing="0" bordercolor="#f0f0f0">          <tbody><tr><td colspan="3"><hr/></td></tr>          <tr bgcolor="#e0e0e0"><td width="35%"><font face="Verdana" size="1" color="#0000bb">Account Number</font></td><td width="10"><font face="Verdana" size="1" color="#0000bb"> : </font></td><td><font face="Verdana" size="1" color="#0000bb">0066941994</font></td></tr>          <tr bgcolor="#f0f0f0"><td width="35%"><font face="Verdana" size="1" color="#0000bb">Name</font></td><td width="10"><font face="Verdana" size="1" color="#0000bb"> : </font></td><td><font face="Verdana" size="1" color="#0000bb">WINSTON SUTANDAR  </font></td></tr>          <tr bgcolor="#e0e0e0"><td width="35%"><font face="Verdana" size="1" color="#0000bb">Period</font></td><td width="10"><font face="Verdana" size="1" color="#0000bb"> : </font></td><td><font face="Verdana" size="1" color="#0000bb">06/10/2017 - 24/10/2017</font></td></tr>          <tr bgcolor="#f0f0f0"><td width="35%"><font face="Verdana" size="1" color="#0000bb">Currency</font></td><td width="10"><font face="Verdana" size="1" color="#0000bb"> : </font></td><td><font face="Verdana" size="1" color="#0000bb">IDR</font></td></tr>          <tr><td colspan="3"><hr/></td></tr>    </tbody></table>  </td></tr><tr>  <td colspan="2">    <table border="1" width="100%" cellpadding="0" cellspacing="0" bordercolor="#ffffff">
<tbody><tr>
<td width="30" bgcolor="e0e0e0"><div align="left">
<font face="Verdana" size="1" color="#0000bb">
<b>Date</b></font></div></td>
<td width="130" bgcolor="e0e0e0"><div align="left">
<font face="Verdana" size="1" color="#0000bb">
<b>Description</b></font></div></td>
<td width="30" bgcolor="e0e0e0"><div align="center">
<font face="Verdana" size="1" color="#0000bb">
<b>Branch</b></font></div></td>
<td width="" bgcolor="e0e0e0" colspan="2"><div align="right">
<font face="Verdana" size="1" color="#0000bb">
<b>Amount</b></font></div></td>
<td width="" bgcolor="e0e0e0"><div align="right">
<font face="Verdana" size="1" color="#0000bb">
<b>Balance</b></font></div></td>
</tr>
<tr>
<td width="30" bgcolor="#e0e0e0"><div align="left">
<font face="verdana" size="1" color="#0000bb">
20/10</font></div></td>
<td width="130" bgcolor="#e0e0e0"><div align="left">
<font face="verdana" size="1" color="#0000bb">
BIAYA ADM         <br/></font></div></td>
<td width="30" bgcolor="#e0e0e0"><div align="center">
<font face="verdana" size="1" color="#0000bb">
0000</font></div></td>
<td width="" bgcolor="#e0e0e0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
5,000.00</font></div></td>
<td width="10" bgcolor="#e0e0e0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
DB</font></div></td>
<td width="" bgcolor="#e0e0e0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
5,000.00</font></div></td>
</tr>
<tr>
<td width="30" bgcolor="#f0f0f0"><div align="left">
<font face="verdana" size="1" color="#0000bb">
PEND </font></div></td>
<td width="130" bgcolor="#f0f0f0"><div align="left">
<font face="verdana" size="1" color="#0000bb">
TRSF E-BANKING CR <br/>10/24 95031       <br/>KREDIT PINTAR     <br/>KEVIN KRISTANTO   <br/></font></div></td>
<td width="30" bgcolor="#f0f0f0"><div align="center">
<font face="verdana" size="1" color="#0000bb">
0000</font></div></td>
<td width="" bgcolor="#f0f0f0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
120,000.00</font></div></td>
<td width="10" bgcolor="#f0f0f0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
CR</font></div></td>
<td width="" bgcolor="#f0f0f0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
125,000.00</font></div></td>
</tr>
<tr>
<td width="30" bgcolor="#e0e0e0"><div align="left">
<font face="verdana" size="1" color="#0000bb">
PEND </font></div></td>
<td width="130" bgcolor="#e0e0e0"><div align="left">
<font face="verdana" size="1" color="#0000bb">
TRSF E-BANKING DB <br/>24/10  WSID:23881 <br/>KEVIN KRISTANTO   <br/></font></div></td>
<td width="30" bgcolor="#e0e0e0"><div align="center">
<font face="verdana" size="1" color="#0000bb">
0000</font></div></td>
<td width="" bgcolor="#e0e0e0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
11,000.00</font></div></td>
<td width="10" bgcolor="#e0e0e0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
DB</font></div></td>
<td width="" bgcolor="#e0e0e0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
114,000.00</font></div></td>
</tr>
<tr>
<td width="30" bgcolor="#f0f0f0"><div align="left">
<font face="verdana" size="1" color="#0000bb">
PEND </font></div></td>
<td width="130" bgcolor="#f0f0f0"><div align="left">
<font face="verdana" size="1" color="#0000bb">
TRSF E-BANKING DB <br/>24/10  WSID:23881 <br/>KEVIN KRISTANTO   <br/></font></div></td>
<td width="30" bgcolor="#f0f0f0"><div align="center">
<font face="verdana" size="1" color="#0000bb">
0000</font></div></td>
<td width="" bgcolor="#f0f0f0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
10,000.00</font></div></td>
<td width="10" bgcolor="#f0f0f0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
DB</font></div></td>
<td width="" bgcolor="#f0f0f0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
104,000.00</font></div></td>
</tr>
<tr>
<td width="30" bgcolor="#e0e0e0"><div align="left">
<font face="verdana" size="1" color="#0000bb">
PEND </font></div></td>
<td width="130" bgcolor="#e0e0e0"><div align="left">
<font face="verdana" size="1" color="#0000bb">
TRSF E-BANKING DB <br/>24/10  WSID:23881 <br/>KEVIN KRISTANTO   <br/></font></div></td>
<td width="30" bgcolor="#e0e0e0"><div align="center">
<font face="verdana" size="1" color="#0000bb">
0000</font></div></td>
<td width="" bgcolor="#e0e0e0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
12,000.00</font></div></td>
<td width="10" bgcolor="#e0e0e0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
DB</font></div></td>
<td width="" bgcolor="#e0e0e0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
92,000.00</font></div></td>
</tr>
<tr>
<td width="30" bgcolor="#f0f0f0"><div align="left">
<font face="verdana" size="1" color="#0000bb">
PEND </font></div></td>
<td width="130" bgcolor="#f0f0f0"><div align="left">
<font face="verdana" size="1" color="#0000bb">
TRSF E-BANKING DB <br/>24/10  WSID:23881 <br/>KEVIN KRISTANTO   <br/></font></div></td>
<td width="30" bgcolor="#f0f0f0"><div align="center">
<font face="verdana" size="1" color="#0000bb">
0000</font></div></td>
<td width="" bgcolor="#f0f0f0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
10,000.00</font></div></td>
<td width="10" bgcolor="#f0f0f0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
DB</font></div></td>
<td width="" bgcolor="#f0f0f0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
82,000.00</font></div></td>
</tr>
<tr>
<td width="30" bgcolor="#e0e0e0"><div align="left">
<font face="verdana" size="1" color="#0000bb">
PEND </font></div></td>
<td width="130" bgcolor="#e0e0e0"><div align="left">
<font face="verdana" size="1" color="#0000bb">
TRSF E-BANKING DB <br/>24/10  WSID:23881 <br/>KEVIN KRISTANTO   <br/></font></div></td>
<td width="30" bgcolor="#e0e0e0"><div align="center">
<font face="verdana" size="1" color="#0000bb">
0000</font></div></td>
<td width="" bgcolor="#e0e0e0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
12,000.00</font></div></td>
<td width="10" bgcolor="#e0e0e0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
DB</font></div></td>
<td width="" bgcolor="#e0e0e0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
70,000.00</font></div></td>
</tr>
<tr>
<td width="30" bgcolor="#f0f0f0"><div align="left">
<font face="verdana" size="1" color="#0000bb">
PEND </font></div></td>
<td width="130" bgcolor="#f0f0f0"><div align="left">
<font face="verdana" size="1" color="#0000bb">
TRSF E-BANKING DB <br/>24/10  WSID:23881 <br/>KEVIN KRISTANTO   <br/></font></div></td>
<td width="30" bgcolor="#f0f0f0"><div align="center">
<font face="verdana" size="1" color="#0000bb">
0000</font></div></td>
<td width="" bgcolor="#f0f0f0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
10,000.00</font></div></td>
<td width="10" bgcolor="#f0f0f0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
DB</font></div></td>
<td width="" bgcolor="#f0f0f0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
60,000.00</font></div></td>
</tr>
<tr>
<td width="30" bgcolor="#e0e0e0"><div align="left">
<font face="verdana" size="1" color="#0000bb">
PEND </font></div></td>
<td width="130" bgcolor="#e0e0e0"><div align="left">
<font face="verdana" size="1" color="#0000bb">
TRSF E-BANKING DB <br/>24/10  WSID:23881 <br/>KEVIN KRISTANTO   <br/></font></div></td>
<td width="30" bgcolor="#e0e0e0"><div align="center">
<font face="verdana" size="1" color="#0000bb">
0000</font></div></td>
<td width="" bgcolor="#e0e0e0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
11,000.00</font></div></td>
<td width="10" bgcolor="#e0e0e0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
DB</font></div></td>
<td width="" bgcolor="#e0e0e0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
49,000.00</font></div></td>
</tr>
<tr>
<td width="30" bgcolor="#f0f0f0"><div align="left">
<font face="verdana" size="1" color="#0000bb">
PEND </font></div></td>
<td width="130" bgcolor="#f0f0f0"><div align="left">
<font face="verdana" size="1" color="#0000bb">
TRSF E-BANKING DB <br/>24/10  WSID:23881 <br/>KEVIN KRISTANTO   <br/></font></div></td>
<td width="30" bgcolor="#f0f0f0"><div align="center">
<font face="verdana" size="1" color="#0000bb">
0000</font></div></td>
<td width="" bgcolor="#f0f0f0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
10,000.00</font></div></td>
<td width="10" bgcolor="#f0f0f0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
DB</font></div></td>
<td width="" bgcolor="#f0f0f0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
39,000.00</font></div></td>
</tr>
<tr>
<td width="30" bgcolor="#e0e0e0"><div align="left">
<font face="verdana" size="1" color="#0000bb">
PEND </font></div></td>
<td width="130" bgcolor="#e0e0e0"><div align="left">
<font face="verdana" size="1" color="#0000bb">
TRSF E-BANKING DB <br/>24/10  WSID:23881 <br/>KEVIN KRISTANTO   <br/></font></div></td>
<td width="30" bgcolor="#e0e0e0"><div align="center">
<font face="verdana" size="1" color="#0000bb">
0000</font></div></td>
<td width="" bgcolor="#e0e0e0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
15,000.00</font></div></td>
<td width="10" bgcolor="#e0e0e0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
DB</font></div></td>
<td width="" bgcolor="#e0e0e0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
24,000.00</font></div></td>
</tr>
<tr>
<td width="30" bgcolor="#f0f0f0"><div align="left">
<font face="verdana" size="1" color="#0000bb">
PEND </font></div></td>
<td width="130" bgcolor="#f0f0f0"><div align="left">
<font face="verdana" size="1" color="#0000bb">
TRSF E-BANKING DB <br/>24/10  WSID:23881 <br/>KEVIN KRISTANTO   <br/></font></div></td>
<td width="30" bgcolor="#f0f0f0"><div align="center">
<font face="verdana" size="1" color="#0000bb">
0000</font></div></td>
<td width="" bgcolor="#f0f0f0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
14,000.00</font></div></td>
<td width="10" bgcolor="#f0f0f0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
DB</font></div></td>
<td width="" bgcolor="#f0f0f0"><div align="right">
<font face="verdana" size="1" color="#0000bb">
10,000.00</font></div></td>
</tr>
</tbody></table>  </td></tr><tr>  <td colspan="2">    <table border="0" width="70%" cellpadding="0" cellspacing="0" bordercolor="#ffffff"> 	 <tbody><tr><td colspan="3"><hr/></td></tr>	 <tr bgcolor="#e0e0e0"><td width="35%"><font face="Verdana" size="1" color="#0000bb">Starting Balance</font></td><td width="10"><font face="Verdana" size="1" color="#0000bb"> : </font></td><td align="right"><font face="Verdana" size="1" color="#0000bb">10,000.00</font></td></tr>	 <tr bgcolor="#f0f0f0"><td width="35%"><font face="Verdana" size="1" color="#0000bb">Total Credits</font></td><td width="10"><font face="Verdana" size="1" color="#0000bb"> : </font></td><td align="right"><font face="Verdana" size="1" color="#0000bb">120,000.00</font></td></tr>	 <tr bgcolor="#e0e0e0"><td width="35%"><font face="Verdana" size="1" color="#0000bb">Total Debits</font></td><td width="10"><font face="Verdana" size="1" color="#0000bb"> : </font></td><td align="right"><font face="Verdana" size="1" color="#0000bb">120,000.00</font></td></tr>	 <tr bgcolor="#f0f0f0"><td width="35%"><font face="Verdana" size="1" color="#0000bb">Ending Balance</font></td><td width="10"><font face="Verdana" size="1" color="#0000bb"> : </font></td><td align="right"><font face="Verdana" size="1" color="#0000bb">10,000.00</font></td></tr>	 <tr><td colspan="3"><hr/></td></tr>    </tbody></table>  </td></tr>
 <form name="iBankForm" method="POST" action="/accountstmt.do"></form>


    </tbody></table>
<table border="0" cellpadding="0" cellspacing="0" width="590" bordercolor="#ffffff">
<tbody><tr height="25">
  <td valign="top">
    <span align="left">
      <font size="1" face="Verdana" color="#f00000"><b> </b> </font>
    </span>
  </td>
</tr>
<tr><td colspan="3"><hr color="#8486ce"/></td></tr>
</tbody></table>
<script language="javascript">
function fncDoubleOut(){
    alert('Your transaction being process. Thank You');
    return false;
}
</script>
`
