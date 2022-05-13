import Keycloak from "keycloak-js";
const keycloak = new Keycloak({
 url: "http://localhost:8080/auth",
 realm: "master",
 clientId: "React-auth-app",
});

export default keycloak;