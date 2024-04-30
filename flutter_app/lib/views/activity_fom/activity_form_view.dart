import 'package:flutter/material.dart';
import 'widgets/activity_form.dart';
import '../../widgets/app_drawer.dart';

class ActivityFormView extends StatelessWidget {
  static const String routeName = '/activity-form';

  const ActivityFormView({super.key});

  @override
  Widget build(BuildContext context) {
    String destinationName =
        ModalRoute.of(context)!.settings.arguments as String;
    return Scaffold(
      appBar: AppBar(
        title: const Text('Ajouter une activité'),
      ),
      drawer: const AppDrawer(),
      body: SingleChildScrollView(
        child: ActivityForm(destinationName: destinationName),
      ),
    );
  }
}
