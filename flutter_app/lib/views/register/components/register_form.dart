import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../../../providers/auth_provider.dart';
import '../../home/home_view.dart';

class RegisterForm extends StatefulWidget {
  const RegisterForm({Key? key}) : super(key: key);

  @override
  _RegisterFormState createState() => _RegisterFormState();
}

class _RegisterFormState extends State<RegisterForm> {
  final _formKey = GlobalKey<FormState>();
  String _firstname = "";
  String _lastname = "";
  String _username = "";
  String _address = "";
  String _email = '';
  String _password = "";
  String _photo = "";

  void _submit() async {
    if (_formKey.currentState!.validate()) {
      _formKey.currentState!.save();
      final authservice = Provider.of<AuthProvider>(context, listen: false);
      bool success = await authservice.Register(_address, _email, _firstname,
          _lastname, _password, _username, _photo);
      if (success) {
        Navigator.pushReplacement(context,
            MaterialPageRoute(builder: (BuildContext context) => HomeView()));
      } else {
        showDialog(
            context: context,
            builder: (context) => AlertDialog(
                  title: Text('Error'),
                  content:
                      Text('Echec lors de l\'inscription, veuillez réessayer'),
                  actions: [
                    TextButton(
                      onPressed: () {
                        Navigator.of(context).pop();
                      },
                      child: Text('OK'),
                    )
                  ],
                ));
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return Form(
        key: _formKey,
        child: Column(children: [
          TextFormField(
            keyboardType: TextInputType.text,
            textInputAction: TextInputAction.next,
            cursorColor: Colors.blue,
            onSaved: (firstname) => _firstname = firstname!,
            decoration: const InputDecoration(
                hintText: "Votre Prénom",
                prefixIcon: Padding(
                  padding: EdgeInsets.all(16.0),
                  child: Icon(Icons.person),
                )),
            validator: (value) {
              if (value == null || value.isEmpty) {
                return 'Veuillez entrer votre prénom';
              }
              return null;
            },
          ),
          Padding(
            padding: const EdgeInsets.symmetric(vertical: 16.0),
            child: TextFormField(
              textInputAction: TextInputAction.done,
              onSaved: (lastname) => _lastname = lastname!,
              decoration: const InputDecoration(
                  hintText: "Votre nom de famille",
                  prefixIcon: Icon(Icons.person)),
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Veuillez entrer votre nom de famille';
                }
                return null;
              },
            ),
          ),
          Padding(
            padding: const EdgeInsets.symmetric(vertical: 16.0),
            child: TextFormField(
              textInputAction: TextInputAction.done,
              onSaved: (username) => _username = username!,
              decoration: const InputDecoration(
                  hintText: "Choisissez un pseudo",
                  prefixIcon: Icon(Icons.person)),
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Veuillez entrer un pseudo';
                }
                return null;
              },
            ),
          ),
          Padding(
            padding: const EdgeInsets.symmetric(vertical: 16.0),
            child: TextFormField(
              textInputAction: TextInputAction.done,
              onSaved: (address) => _address = address!,
              decoration: const InputDecoration(
                  hintText: "Votre adresse", prefixIcon: Icon(Icons.place)),
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Veuillez entrer votre adresse postale';
                }
                return null;
              },
            ),
          ),
          Padding(
            padding: const EdgeInsets.symmetric(vertical: 16.0),
            child: TextFormField(
              textInputAction: TextInputAction.done,
              onSaved: (email) => _email = email!,
              decoration: const InputDecoration(
                  hintText: "Votre adresse mail", prefixIcon: Icon(Icons.mail)),
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Veuillez entrer adresse mail';
                }
                return null;
              },
            ),
          ),
          Padding(
            padding: const EdgeInsets.symmetric(vertical: 16.0),
            child: TextFormField(
              textInputAction: TextInputAction.done,
              obscureText: true,
              cursorColor: Colors.blue,
              onSaved: (password) => _password = password!,
              decoration: const InputDecoration(
                hintText: "Votre mot de passe",
                prefixIcon: Padding(
                  padding: EdgeInsets.all(16.0),
                  child: Icon(Icons.lock),
                ),
              ),
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Veuillez entrer votre mot de passe';
                }
                return null;
              },
            ),
          ),
          const SizedBox(height: 16.0),
          ElevatedButton(
            onPressed: _submit,
            child: Text(
              "S\'inscrire".toUpperCase(),
            ),
          ),
          const SizedBox(height: 16.0),
        ]));
  }
}
