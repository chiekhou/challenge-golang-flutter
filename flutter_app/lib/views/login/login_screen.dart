import 'package:flutter/material.dart';
import 'package:flutter_app/views/login/components/login_form.dart';
import 'package:provider/provider.dart';
import '../../providers/auth_provider.dart';

class LoginScreen extends StatelessWidget {
  static const String routeName = '/';
  const LoginScreen({super.key});

  @override
  Widget build(BuildContext context) {
    Size size = MediaQuery.of(context).size;
    TextEditingController usernameController = TextEditingController();
    TextEditingController passwordController = TextEditingController();
    AuthProvider authProvider = Provider.of<AuthProvider>(context);

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
/*            Padding(
              padding: const EdgeInsets.symmetric(horizontal: 20),
              child: Column(
                children: [
                  TextField(
                    controller: usernameController,
                    decoration: InputDecoration(labelText: "Username"),
                  ),
                  TextField(
                    controller: passwordController,
                    decoration: InputDecoration(labelText: "Password"),
                    obscureText: true,
                  ),
                ],
              ),
            ),*/
            SizedBox(height: size.height * 0.03),
            LoginForm(),

/*
            ElevatedButton(
              onPressed: () {
                authProvider.login(
                  usernameController.text,
                  passwordController.text,
                );

                if (authProvider.isAuthenticated) {
                  Navigator.pushReplacementNamed(
                    context,
                    authProvider.isAdmin ? '/admin_dashboard' : '/',
                  );
                } else {
                  // Afficher un message d'erreur si la connexion échoue
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(content: Text('Invalid username or password')),
                  );
                }
              },
              child: Text('Login'),
            ),
*/
            Spacer(flex: 3),
          ],
        ),
      ),
    );
  }
}
