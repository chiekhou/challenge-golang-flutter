import 'package:flutter/material.dart';
import 'package:flutter_app/providers/auth_provider.dart';
import 'package:flutter_app/views/home/home_view.dart';
import 'package:provider/provider.dart';

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
      final authService = Provider.of<AuthProvider>(context, listen: false);
      bool success = await authService.Register(
        address: _address,
        email: _email,
        firstname: _firstname,
        lastname: _lastname,
        password: _password,
        username: _username,
        photo: _photo,
      );
      if (success) {
        Navigator.pushReplacement(
          context,
          MaterialPageRoute(builder: (BuildContext context) => HomeView()),
        );
      } else {
        showDialog(
          context: context,
          builder: (context) => AlertDialog(
            title: Text('Erreur'),
            content: Text('Échec lors de l\'inscription, veuillez réessayer'),
            actions: [
              TextButton(
                onPressed: () {
                  Navigator.of(context).pop();
                },
                child: Text('OK'),
              )
            ],
          ),
        );
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return Form(
      key: _formKey,
      child: Column(
        children: [
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
              ),
            ),
            validator: (value) {
              if (value == null || value.isEmpty) {
                return 'Veuillez entrer votre prénom';
              }
              return null;
            },
          ),
          // Ajoutez ici les autres champs du formulaire
          ElevatedButton(
            onPressed: _submit,
            child: Text("S'inscrire"),
          ),
        ],
      ),
    );
  }
}
