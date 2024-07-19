import 'package:flutter/material.dart';
import 'package:flutter_app/providers/group_voyage_provider.dart';
import 'package:provider/provider.dart';

class AddGroupFormAdmin extends StatefulWidget {
  static const routeName = '/admin/add_group';

  const AddGroupFormAdmin({super.key});

  @override
  _AddGroupFormState createState() => _AddGroupFormState();
}

class _AddGroupFormState extends State<AddGroupFormAdmin> {
  final _formKey = GlobalKey<FormState>();
  double _budget = 0.0;
  String _nom = "";
  String _photo = "";
  int? _voyageId;

  bool _isLoading = false;

  void _submitForm() async {
    if (_formKey.currentState!.validate()) {
      setState(() {
        _isLoading = true;
      });

      String nom = _nom.trim();
      double budget = _budget;
      String photo = _photo.trim();

      final groupService =
          Provider.of<GroupVoyageProvider>(context, listen: false);

      bool success = await groupService.CreateGroup(budget, nom, _voyageId);

      setState(() {
        _isLoading = false;
      });

      if (success) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(
            content: Text('Groupe ajouté avec succès!'),
            backgroundColor: Colors.green,
          ),
        );
      } else {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(
            content: Text('Erreur lors de l\'ajout du groupe'),
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
        title: const Text('Ajouter un Groupe'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Form(
          key: _formKey,
          child: Card(
            elevation: 8.0,
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(10.0),
            ),
            child: Padding(
              padding: const EdgeInsets.all(16.0),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: <Widget>[
                  TextFormField(
                    decoration:
                        const InputDecoration(labelText: 'Nom du groupe'),
                    onChanged: (value) {
                      setState(() {
                        _nom = value;
                      });
                    },
                    validator: (value) {
                      if (value == null || value.isEmpty) {
                        return 'Entrez le nom du groupe';
                      }
                      return null;
                    },
                  ),
                  TextFormField(
                    decoration: const InputDecoration(labelText: 'Budget'),
                    keyboardType: TextInputType.number,
                    onChanged: (value) {
                      setState(() {
                        _budget = double.tryParse(value) ?? 0.0;
                      });
                    },
                    validator: (value) {
                      if (value == null || value.isEmpty) {
                        return 'Entrez le budget';
                      }
                      if (double.tryParse(value) == null) {
                        return 'Entrez un nombre valide';
                      }
                      return null;
                    },
                  ),
                  // TextFormField(
                  //   decoration: const InputDecoration(labelText: 'Photo (URL)'),
                  //   onChanged: (value) {
                  //     setState(() {
                  //       _photo = value;
                  //     });
                  //   },
                  //   validator: (value) {
                  //     return null;
                  //   },
                  // ),
                  const SizedBox(height: 16),
                  _isLoading
                      ? const Center(child: CircularProgressIndicator())
                      : ElevatedButton(
                          onPressed: _submitForm,
                          child: const Text('Ajouter'),
                        ),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}
