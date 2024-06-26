import "package:flutter/material.dart";
import "package:flutter_app/views/login/components/login_form.dart";

class LoginScreen extends StatelessWidget{
  static const String routeName = '/';
  const LoginScreen({super.key});
  @override
  Widget build(BuildContext context) {
    Size size = MediaQuery.of(context).size;
    return Scaffold(
      body: Container(
        width: double.infinity,
        height: size.height,
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Spacer(flex: 2),
            Text(
              "LOGIN",
              style: TextStyle(
                fontWeight: FontWeight.bold,
                fontSize: 26,
              ),
            ),
            SizedBox(height: size.height * 0.03),
            LoginForm(),
            SizedBox(height: size.height * 0.03),
            Spacer(flex: 3)
          ],
        )
      )
    );
    throw UnimplementedError();
  }
}
