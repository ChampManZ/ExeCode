import { CognitoUserPool } from "amazon-cognito-identity-js";

const poolData = {
    UserPoolId: "us-east-1_wtxlFGAcv",
    ClientId: "5ujg8c9bbiihttb1gdsijj48t0", 
}

export default new CognitoUserPool(poolData);
