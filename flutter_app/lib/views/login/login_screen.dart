import 'package:flutter/material.dart';
import 'package:flutter_app/views/login/components/login_form.dart';

class LoginScreen extends StatelessWidget {
  static const String routeName = '/';
  const LoginScreen({super.key});
  @override
  Widget build(BuildContext context) {
    Size size = MediaQuery.of(context).size;
    return Scaffold(
      body: Stack(
        children: [
          Container(
            width: double.infinity,
            height: size.height,
            decoration: BoxDecoration(
              image: DecorationImage(
                image: AssetImage('assets/images/background.jpg'),
                fit: BoxFit.cover,
              ),
            ),
          ),
          Container(
            width: double.infinity,
            height: size.height,
            color: Colors.white.withOpacity(0.5),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Spacer(flex: 2),
                Text(
                  "CONNEXION",
                  style: TextStyle(
                    fontWeight: FontWeight.bold,
                    fontSize: 32,
                    color: Colors.black,
                  ),
                ),
                SizedBox(height: size.height * 0.03),
                Padding(
                  padding: EdgeInsets.symmetric(horizontal: 30),
                  child: LoginForm(),
                ),
                SizedBox(height: size.height * 0.03),
                Spacer(flex: 3)
              ],
            ),
          ),
        ],
      ),
    );
  }
}