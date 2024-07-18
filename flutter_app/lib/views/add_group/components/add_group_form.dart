import 'dart:ui';

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_app/views/group_voyage/groupeVoyage_screen.dart';
import 'package:provider/provider.dart';

import '../../../providers/group_voyage_provider.dart';

class AddGroupForm extends StatefulWidget{
  const AddGroupForm({Key? key}): super(key: key);

  @override
  _AddGroupFormState createState() => _AddGroupFormState();
}

class _AddGroupFormState extends State<AddGroupForm>{
  final _formkey = GlobalKey<FormState>();
  double _budget = 0.0;
  String _nom = "";
  String _photo = "";

  void _submit()async{
    if(_formkey.currentState!.validate()){
      _formkey.currentState!.save();
      final groupeVoyageService = Provider.of<GroupVoyageProvider>(
        context, listen: false
      );
      bool success = await groupeVoyageService.CreateGroup(_budget, _nom);
      if(success){
        Navigator.pushReplacement(context,
        MaterialPageRoute(builder: (BuildContext context) =>GroupevoyageScreen()));
      }else{
        showDialog(
            context: context,
            builder: (context)=>AlertDialog(
              title: Text('Error'),
              content: Text('Erreur sur l\'ajout de groupe'),
              actions: [
                TextButton(onPressed: (){
                  Navigator.of(context).pop();
                },
                  child: Text('ok')
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
        children: [
          TextFormField(
            keyboardType: TextInputType.name,
            textInputAction: TextInputAction.next,
            onSaved: (nom)=> _nom = nom!,
            decoration: const InputDecoration(
              hintText: 'Nom de votre groupe',
                border: OutlineInputBorder(
                    borderRadius: BorderRadius.all(Radius.circular(30.0))
                )
            ),
            validator: (value) {
              if(value == null || value.isEmpty){
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
              onSaved: (budget)=> _budget = double.parse(budget!),
              decoration: const InputDecoration(
                hintText: 'Votre budget',
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.all(Radius.circular(30.0))
                )
              ),
            )
          ),
          const SizedBox(height: 32.0),
          ElevatedButton(
              onPressed: _submit,
              child: Text('Créer le groupe de voyage'.toUpperCase()),
          )
        ],
      )
    );
  }
}