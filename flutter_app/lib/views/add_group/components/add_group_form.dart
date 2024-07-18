import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../../../providers/group_voyage_provider.dart';
import '../../group_voyage/groupeVoyage_screen.dart';
import 'voyage_dropdown_menu.dart';

class AddGroupForm extends StatefulWidget {
  const AddGroupForm({Key? key}) : super(key: key);

  @override
  _AddGroupFormState createState() => _AddGroupFormState();
}

class _AddGroupFormState extends State<AddGroupForm> {
  final _formKey = GlobalKey<FormState>();
  double _budget = 0.0;
  String _nom = "";
  int? _voyageId;

  void _submit() async {
    if (_formKey.currentState!.validate()) {
      _formKey.currentState!.save();
      final groupVoyageProvider =
          Provider.of<GroupVoyageProvider>(context, listen: false);
      bool success =
          await groupVoyageProvider.CreateGroup(_budget, _nom, _voyageId);
      if (success) {
        Navigator.pushReplacement(
          context,
          MaterialPageRoute(
              builder: (BuildContext context) => GroupevoyageScreen()),
        );
      } else {
        showDialog(
          context: context,
          builder: (context) => AlertDialog(
            title: Text('Error'),
            content: Text('Erreur sur l\'ajout de groupe'),
            actions: [
              TextButton(
                onPressed: () {
                  Navigator.of(context).pop();
                },
                child: Text('OK'),
              ),
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
            keyboardType: TextInputType.name,
            textInputAction: TextInputAction.next,
            onSaved: (nom) => _nom = nom!,
            decoration: const InputDecoration(
              hintText: 'Nom de votre groupe',
              border: OutlineInputBorder(
                borderRadius: BorderRadius.all(Radius.circular(30.0)),
              ),
            ),
            validator: (value) {
              if (value == null || value.isEmpty) {
                return 'Veuillez entrer un nom de groupe';
              }
              return null;
            },
          ),
          Padding(
            padding: const EdgeInsets.symmetric(vertical: 20.0),
            child: TextFormField(
              keyboardType: TextInputType.number,
              textInputAction: TextInputAction.done,
              onSaved: (budget) => _budget = double.parse(budget!),
              decoration: const InputDecoration(
                hintText: 'Votre budget',
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.all(Radius.circular(30.0)),
                ),
              ),
            ),
          ),
          const SizedBox(height: 16.0),
          VoyageDropdownMenu(
            onSelectedVoyageChanged: (int? voyageId) {
              _voyageId = voyageId;
            },
          ),
          const SizedBox(height: 32.0),
          ElevatedButton(
            onPressed: _submit,
            child: Text('Créer le groupe de voyage'.toUpperCase()),
          ),
        ],
      ),
    );
  }
}
