package middlewares

// func Authentication() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		verifyToken, err := helpers.VerifyToken(ctx)
// 		_ = verifyToken

// 		if err != nil {
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error":   "Unauthenticated",
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 		ctx.Set("userData", verifyToken)
// 		ctx.Next()
// 	}
// }
