import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../../providers/auth_provider.dart';
import '../home/home_view.dart';

class RegisterScreen extends StatefulWidget {
  static const String routeName = '/register';

  const RegisterScreen({Key? key}) : super(key: key);

  @override
  _RegisterScreenState createState() => _RegisterScreenState();
}

class _RegisterScreenState extends State<RegisterScreen> {
  int _currentStep = 0;
  final _formKey = GlobalKey<FormState>();
  String _firstname = "";
  String _lastname = "";
  String _username = "";
  String _address = "";
  String _email = '';
  String _password = "";
  String _photo = ""; // Valeur par défaut pour le paramètre photo

  void _nextStep() {
    if (_currentStep < 1) {
      setState(() {
        _currentStep++;
      });
    }
  }

  void _previousStep() {
    if (_currentStep > 0) {
      setState(() {
        _currentStep--;
      });
    }
  }

  Future<void> submitForm() async {
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
            content:
                Text('Échec lors de l\'inscription, veuillez réessayer'),
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
    return Scaffold(
      appBar: AppBar(
        title: Text('Inscription'),
        centerTitle: true,
      ),
      body: Container(
        padding: const EdgeInsets.all(20),
        child: Form(
          key: _formKey,
          child: Stepper(
            type: StepperType.horizontal,
            currentStep: _currentStep,
            onStepCancel: () => _currentStep == 0 ? null : _previousStep(),
            onStepContinue: () {
              bool isLastStep = (_currentStep == getSteps().length - 1);
              if (isLastStep) {
                submitForm();
              } else {
                _nextStep();
              }
            },
            onStepTapped: (step) => setState(() {
              _currentStep = step;
            }),
            steps: getSteps(),
          ),
        ),
      ),
    );
  }

  List<Step> getSteps() {
    return <Step>[
      Step(
        state: _currentStep > 0 ? StepState.complete : StepState.indexed,
        isActive: _currentStep >= 0,
        title: Text(''),
        content: Column(
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
            Padding(
              padding: const EdgeInsets.symmetric(vertical: 16.0),
              child: TextFormField(
                textInputAction: TextInputAction.done,
                onSaved: (lastname) => _lastname = lastname!,
                decoration: const InputDecoration(
                  hintText: "Votre nom de famille",
                  prefixIcon: Icon(Icons.person),
                ),
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
                onSaved: (address) => _address = address!,
                decoration: const InputDecoration(
                  hintText: "Votre adresse",
                  prefixIcon: Icon(Icons.location_city),
                ),
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return 'Veuillez entrer votre Adresse';
                  }
                  return null;
                },
              ),
            ),
          ],
        ),
      ),
      Step(
        state: _currentStep > 1 ? StepState.complete : StepState.indexed,
        isActive: _currentStep >= 1,
        title: Text(''),
        content: Column(
          children: [
            TextFormField(
              keyboardType: TextInputType.text,
              onSaved: (username) => _username = username!,
              decoration: const InputDecoration(
                hintText: 'Entrez un username',
              ),
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Veuillez entrer votre Username';
                }
                return null;
              },
            ),
            Padding(
              padding: EdgeInsets.all(10.0),
              child: TextFormField(
                textInputAction: TextInputAction.done,
                onSaved: (email) => _email = email!,
                decoration: const InputDecoration(
                  hintText: 'Entrez votre email',
                ),
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return 'Veuillez entrer votre email';
                  }
                  return null;
                },
              ),
            ),
            Padding(
              padding: EdgeInsets.all(10.0),
              child: TextFormField(
                textInputAction: TextInputAction.done,
                onSaved: (password) => _password = password!,
                decoration: const InputDecoration(
                  hintText: 'Entrez un mot de passe',
                ),
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return 'Veuillez entrer votre mot de passe';
                  }
                  return null;
                },
              ),
            ),
          ],
        ),
      ),
    ];
  }
}
