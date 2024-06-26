import 'package:flutter/material.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter_app/views/register/components/register_form.dart';

class RegisterScreen extends StatelessWidget{
  static const String routeName = '/register';
  const RegisterScreen({super.key});
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
                  "INSCRIPTION",
                  style: TextStyle(
                    fontWeight: FontWeight.bold,
                    fontSize: 26,
                  ),
                ),
                SizedBox(height: size.height * 0.03),
                RegisterForm(),
                SizedBox(height: size.height * 0.03),
                Spacer(flex: 3)
              ],
            )
        )
    );
    throw UnimplementedError();
  }
}