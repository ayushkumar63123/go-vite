(window.webpackJsonp=window.webpackJsonp||[]).push([[116],{714:function(e,a,t){"use strict";t.r(a);var s=t(1),i=Object(s.a)({},(function(){var e=this,a=e.$createElement,t=e._self._c||a;return t("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[t("h1",{attrs:{id:"migrating-from-solidity-0-4"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#migrating-from-solidity-0-4"}},[e._v("#")]),e._v(" Migrating from Solidity++ 0.4")]),e._v(" "),t("p",[e._v("Solidity++ 0.8 is syntactically closer to Solidity than Solidity++ 0.4, which is more friendly to developers from Solidity.")]),e._v(" "),t("p",[e._v("Some counter-intuitive syntex and keywords are removed in this version.")]),e._v(" "),t("h2",{attrs:{id:"solidity-0-8-0-breaking-changes"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#solidity-0-8-0-breaking-changes"}},[e._v("#")]),e._v(" Solidity++ 0.8.0 Breaking Changes")]),e._v(" "),t("ul",[t("li",[t("p",[e._v("All Solidity breaking changes since v0.5.0 to v0.8.0 are applied to Solidity++ 0.8.")])]),e._v(" "),t("li",[t("p",[t("code",[e._v("onMessage")]),e._v(" keyword is deprecated since 0.8.0. Use function declarations instead ("),t("code",[e._v("function external async")]),e._v(").")])]),e._v(" "),t("li",[t("p",[t("code",[e._v("message")]),e._v(" keyword is deprecated since 0.8.0. Use function types instead.")])]),e._v(" "),t("li",[t("p",[t("code",[e._v("send")]),e._v(" keyword is deprecated since 0.8.0. Use function calls instead.")])]),e._v(" "),t("li",[t("p",[t("code",[e._v("getter")]),e._v(" keyword is deprecated since 0.8.0. Use offchain function instead ("),t("code",[e._v("function offchain")]),e._v(").")])]),e._v(" "),t("li",[t("p",[e._v("Inline assembly and Yul are available in this version.")])]),e._v(" "),t("li",[t("p",[t("code",[e._v("keccak256")]),e._v(" and "),t("code",[e._v("sha256")]),e._v(" are available in this version.")])])]),e._v(" "),t("h2",{attrs:{id:"how-to-update-your-code"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#how-to-update-your-code"}},[e._v("#")]),e._v(" How to Update Your Code")]),e._v(" "),t("p",[e._v("Let's start with an example:")]),e._v(" "),t("tm-code-block",{staticClass:"codeblock",attrs:{language:"javascript",base64:"cHJhZ21hIHNvbGlkaXR5cHAgXjAuNC4zOwoKY29udHJhY3QgQSB7CiAgICBtZXNzYWdlIHN1bSh1aW50IHJlc3VsdCk7CiAgICAKICAgIG9uTWVzc2FnZSBhZGQodWludCBhLCB1aW50IGIpIHsKICAgICAgICB1aW50IHJlc3VsdCA9IGEgKyBiOwogICAgICAgIGFkZHJlc3Mgc2VuZGVyID0gbXNnLnNlbmRlcjsKICAgICAgICBzZW5kKHNlbmRlciwgc3VtKHJlc3VsdCkpOwogICB9Cn0KCmNvbnRyYWN0IEIgewogICAgYWRkcmVzcyBhZGRyQTsKICAgIHVpbnQgdG90YWw7CiAgICBtZXNzYWdlIGFkZCh1aW50IGEsIHVpbnQgYik7CgogICAgY29uc3RydWN0b3IgKGFkZHJlc3MgYWRkcikgewogICAgICAgIGFkZHJBID0gYWRkcjsKICAgIH0KCiAgICBvbk1lc3NhZ2UgaW52b2tlKHVpbnQgYSwgdWludCBiKSB7CiAgICAgICAgc2VuZChhZGRyQSwgYWRkKGEsIGIpKTsKICAgIH0KCiAgICBvbk1lc3NhZ2Ugc3VtKHVpbnQgcmVzdWx0KSB7CiAgICAgICAgdG90YWwgKz0gcmVzdWx0OwogICAgfQoKICAgIGdldHRlciB0b3RhbCgpIHJldHVybnModWludCkgewogICAgICAgIHJldHVybiB0b3RhbDsKICAgIH0KfQo="}}),e._v(" "),t("p",[e._v("In the above code, contract A declares a message listener "),t("code",[e._v("add(uint a, uint b)")]),e._v(".")]),e._v(" "),t("p",[e._v("contract B declares "),t("code",[e._v("add")]),e._v(" message which has the same signature to "),t("code",[e._v("add")]),e._v(" message listener in contract A.")]),e._v(" "),t("p",[e._v("Contract B declares a message listener "),t("code",[e._v("invoke")]),e._v(" as the entry to the contract. When "),t("code",[e._v("B.invoke()")]),e._v(" is called, contract B sends a "),t("code",[e._v("add")]),e._v(" message to contract A to initiate an asynchronous message call.")]),e._v(" "),t("p",[e._v("When contract A responds to the message call, it sends a "),t("code",[e._v("sum")]),e._v(" message back to contract B to return data asynchronously.")]),e._v(" "),t("p",[e._v("Contract B also declares a message listener "),t("code",[e._v("sum(uint result)")]),e._v(" as a "),t("em",[e._v("'callback function'")]),e._v(" to handle the returned message from contract A.")]),e._v(" "),t("p",[e._v("Since Solidity++ 0.8.0, no explicit "),t("code",[e._v("onMessage")]),e._v(" declaration as a callback is required. The compiler is smart enough to generate callbacks automatically. The contract code is simplified and optimized significantly by async/await syntactic sugar.")]),e._v(" "),t("p",[e._v("The migrated code is as follows:")]),e._v(" "),t("tm-code-block",{staticClass:"codeblock",attrs:{language:"javascript",base64:"Ly8gU1BEWC1MaWNlbnNlLUlkZW50aWZpZXI6IEdQTC0zLjAKcHJhZ21hIHNvbGlkaXR5cHAgJmd0Oz0wLjguMCAmbHQ7MC45LjA7Cgpjb250cmFjdCBBIHsKICAgIC8vIHRoZSBvbk1lc3NhZ2UgaXMgcmVwbGFjZWQgd2l0aCBhbiBhc3luYyBmdW5jdGlvbgogICAgZnVuY3Rpb24gYWRkKHVpbnQgYSwgdWludCBiKSBleHRlcm5hbCBhc3luYyByZXR1cm5zKHVpbnQpIHsKICAgICAgICByZXR1cm4gYSArIGI7CiAgICB9Cn0KCmNvbnRyYWN0IEIgewogICAgQSBjb250cmFjdEE7CiAgICB1aW50IHB1YmxpYyB0b3RhbDsgIC8vIGFuIG9mZmNoYWluIGdldHRlciBmdW5jdGlvbiB3aWxsIGJlIGF1dG8tZ2VuZXJhdGVkCgogICAgY29uc3RydWN0b3IgKGFkZHJlc3MgYWRkcikgewogICAgICAgIGNvbnRyYWN0QSA9IEEoYWRkcik7CiAgICB9CgogICAgLy8gdGhlIG9uTWVzc2FnZSBpcyByZXBsYWNlZCB3aXRoIGFuIGFzeW5jIGZ1bmN0aW9uCiAgICBmdW5jdGlvbiBpbnZva2UodWludCBhLCB1aW50IGIpIGV4dGVybmFsIGFzeW5jIHsKICAgICAgICAvLyB1c2UgYXdhaXQgb3BlcmF0b3IgdG8gZ2V0IHJldHVybiBkYXRhCiAgICAgICAgdG90YWwgKz0gYXdhaXQgY29udHJhY3RBLmFkZChhLCBiKTsKICAgIH0KfQo="}})],1)}),[],!1,null,null,null);a.default=i.exports}}]);