# okta-go-inline-hook

This is a sample okta inline saml hook written in golang as a aws lambda function.  https://developer.okta.com/docs/concepts/inline-hooks/.  

1) You will need an okta org.  https://developer.okta.com/signup/

2) Next set up a an inline hook in your org. https://help.okta.com/en/prod/Content/Topics/automation-hooks/add-inline-hooks.htm

3) This inline hook will let you control who is allowed to register to your org https://developer.okta.com/docs/reference/registration-hook/. 

4) The inline hook will reach out to your REST based API (webservice) with an endpoint of /registration.

5) The sample app needs to be hosted in a publicly accessible address with an ssl/https enabled port 
  - This sample includes a makefile and toml file set up specifically for https://www.netlify.com/.

6) In your API, it will need to take an incoming request and in the return body you will need to send it a json string that the inline hook will recognize 