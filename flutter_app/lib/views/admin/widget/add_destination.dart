import 'package:flutter/material.dart';
import 'package:flutter_app/providers/destination_provider.dart';
import 'package:flutter_app/views/admin/home_dashboard.dart';
import 'package:provider/provider.dart';

class AddDestinationForm extends StatefulWidget {
  static const routeName = '/admin/add_destination';

  @override
  _AddDestinationFormState createState() => _AddDestinationFormState();
}

class _AddDestinationFormState extends State<AddDestinationForm> {
  final _formKey = GlobalKey<FormState>();

  final TextEditingController _destinationNameController =
      TextEditingController();
  final TextEditingController _urlImageController = TextEditingController();

  bool _isLoading = false;

  void _submitForm() async {
    if (_formKey.currentState!.validate()) {
      setState(() {
        _isLoading = true;
      });

      String name = _destinationNameController.text.trim();
      String image = _urlImageController.text.trim();

      final destinationService =
          Provider.of<DestinationProvider>(context, listen: false);

      bool success = await destinationService.CreateDestination(name, image);

      setState(() {
        _isLoading = false;
      });

      if (success) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(
            content: Text('Destination Crée !'),
            backgroundColor: Colors.green,
          ),
        );
        Navigator.pushNamed(context, DashboardHomeScreen.routeName);
        ;
      } else {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(
            content: Text('Erreur lors de la création de destination'),
            backgroundColor: Colors.red,
          ),
        );
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Ajouter une destination'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: SingleChildScrollView(
          child: Form(
            key: _formKey,
            child: Card(
              elevation: 4.0,
              child: Padding(
                padding: const EdgeInsets.all(16.0),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: <Widget>[
                    const Text(
                      'Détails de la destination',
                      style:
                          TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
                    ),
                    const SizedBox(height: 16),
                    TextFormField(
                      controller: _destinationNameController,
                      decoration: const InputDecoration(
                          labelText: 'Nom de destination'),
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return 'Entrez une destination';
                        }
                        return null;
                      },
                    ),
                    const SizedBox(height: 16),
                    TextFormField(
                      controller: _urlImageController,
                      decoration: const InputDecoration(labelText: 'Image'),
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return 'Entrez l\'url de l\'image';
                        }
                        return null;
                      },
                    ),
                    const SizedBox(height: 32),
                    _isLoading
                        ? Center(child: CircularProgressIndicator())
                        : SizedBox(
                            width: double.infinity,
                            child: ElevatedButton(
                              onPressed: _submitForm,
                              child: const Text('Ajouter'),
                            ),
                          ),
                  ],
                ),
              ),
            ),
          ),
        ),
      ),
    );
  }
}
