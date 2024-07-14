import 'package:flutter/material.dart';
import 'package:flutter_app/providers/auth_provider.dart';
import 'package:flutter_app/views/home/home_view.dart';
import 'package:flutter_app/views/register/register_screen.dart';
import 'package:provider/provider.dart';

class LoginForm extends StatefulWidget {
  const LoginForm({Key? key}) : super(key: key);

  @override
  _LoginFormState createState() => _LoginFormState();
}

class _LoginFormState extends State<LoginForm> {
  final _formkey = GlobalKey<FormState>();
  String _email = "";
  String _password = "";

  void _submit() async{
    if(_formkey.currentState!.validate()){
      _formkey.currentState!.save();
      final authservice = Provider.of<AuthProvider>(context, listen: false);
      bool success = await authservice.login(_email, _password);
      if(success){

        Navigator.pushReplacement(context,
            MaterialPageRoute(builder: (BuildContext context) => HomeView()));
      }else{
        showDialog(
            context: context,
            builder: (context) => AlertDialog(
          title: Text('Error'),
          content: Text('Connexion failed, please try again'),
          actions: [
            TextButton(onPressed: (){
              Navigator.of(context).pop();
            },
            child: Text('OK'),
            )
          ],
        )
        );
      }
    }
  }

  @override
  Widget build(BuildContext context){
    return Form(
      key: _formkey,
        child: Column(
            children:[
              TextFormField(
                keyboardType: TextInputType.emailAddress,
                textInputAction: TextInputAction.next,
                cursorColor: Colors.blue,
                onSaved: (email)=> _email = email!,
                decoration: const InputDecoration(
                    hintText: "Votre email",
                    prefixIcon: Padding(
                      padding: EdgeInsets.all(16.0),
                      child: Icon(Icons.person),
                    )
                ),
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return 'Veuillez entrer votre email';
                  }
                  return null;
                },
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
                  "Connexion".toUpperCase(),
                ),
              ),
              const SizedBox(height: 16.0),
              ListTile(
                onTap: (){
                  Navigator.pushReplacement(
                      context,
                      MaterialPageRoute(builder: (BuildContext context) => RegisterScreen()));
                },
                title: Text('Créer un compte')
              ),
            ]
        )
    );
  }
}


