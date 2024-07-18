import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_app/views/add_group/components/add_group_form.dart';

class AddGroupScreen extends StatelessWidget {
  static const String routeName = '/add_group';
  const AddGroupScreen({super.key});

  @override
  Widget build(BuildContext context) {
    Size size = MediaQuery.of(context).size;
    return Scaffold(
      appBar: AppBar(),
      body: Container(
        width: double.infinity,
        height: size.height,
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Spacer(flex: 2),
            const Text('Créer un groupe',
                style: TextStyle(
                  fontWeight: FontWeight.bold,
                  fontSize: 26,
                )),
            SizedBox(height: size.height * 0.03),
            AddGroupForm(),
            SizedBox(height: size.height * 0.03),
            Spacer(flex: 3)
          ],
        ),
      ),
    );
  }
}
