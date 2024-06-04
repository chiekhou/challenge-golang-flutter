import 'package:flutter/material.dart';
import 'package:flutter_app/views/activity_fom/activity_form_view.dart';
import 'package:flutter_app/widgets/app_drawer.dart';
import 'package:provider/provider.dart';
import '../../models/destination_model.dart';
import '../../models/activity_model.dart';
import '../../models/voyage_model.dart';
import '../../providers/destination_provider.dart';
import '../../providers/voyage_provider.dart';
import '../home/home_view.dart';
import 'widgets/voyage_activity_list.dart';
import 'widgets/activity_list.dart';
import 'widgets/voyage_overview.dart';

class DestinationView extends StatefulWidget {
  static const String routeName = '/destination';

  const DestinationView({super.key});

  @override
  State<DestinationView> createState() => _DestinationState();
}

class _DestinationState extends State<DestinationView> {
  late Voyage myvoyage;
  late int index;

  @override
  void initState() {
    super.initState();

    index = 0;
    myvoyage = Voyage(
      activities: [],
      date: null,
      destination: '',
    );
  }

  double get amount {
    return myvoyage.activities.fold(0.0, (prev, element) {
      return prev + element.price;
    });
  }

  void setDate() {
    showDatePicker(
      context: context,
      initialDate: DateTime.now().add(const Duration(days: 1)).toUtc(),
      firstDate: DateTime.now().toUtc(),
      lastDate: DateTime(2030),
    ).then((newDate) {
      if (newDate != null) {
        setState(() {
          myvoyage.date = newDate.toUtc();
        });
      }
    });
  }

  void switchIndex(newIndex) {
    setState(() {
      index = newIndex;
    });
  }

  void toggleActivity(Activity activity) {
    setState(() {
      myvoyage.activities.contains(activity)
          ? myvoyage.activities.remove(activity)
          : myvoyage.activities.add(activity);
    });
  }

  void deleteVoyageActivity(Activity activity) {
    setState(() {
      myvoyage.activities.remove(activity);
    });
  }

  void saveVoyage(String destinationName) async {
    final result = await showDialog(
      context: context,
      builder: (context) {
        return SimpleDialog(
          title: const Text('Voulez-vous sauvegarder ?'),
          contentPadding: const EdgeInsets.all(20),
          children: <Widget>[
            Row(
              mainAxisAlignment: MainAxisAlignment.end,
              children: <Widget>[
                ElevatedButton(
                  child: const Text('Annuler'),
                  onPressed: () {
                    Navigator.pop(context, 'cancel');
                  },
                ),
                const SizedBox(
                  width: 20,
                ),
                ElevatedButton(
                  style: ElevatedButton.styleFrom(
                      backgroundColor: Theme.of(context).primaryColor),
                  onPressed: () {
                    Navigator.pop(context, 'save');
                  },
                  child: const Text(
                    'Sauvegarder',
                    style: TextStyle(color: Colors.white),
                  ),
                ),
              ],
            ),
          ],
        );
      },
    );

    print('Résultat de la boîte de dialogue: $result');

    if (myvoyage.date == null) {
      print('Date du voyage non définie');
      if (mounted) {
        showDialog(
          context: context,
          builder: (context) => AlertDialog(
            title: const Text('Attention !'),
            content: const Text('Vous n\'avez pas entré de date'),
            actions: <Widget>[
              TextButton(
                child: const Text('Ok'),
                onPressed: () => Navigator.pop(context),
              ),
            ],
          ),
        );
      }
    } else if (result == 'save') {
      if (mounted) {
        myvoyage.destination = destinationName;
        print('Sauvegarde du voyage avant ajout au fournisseur: $myvoyage');
        Provider.of<VoyageProvider>(context, listen: false).addVoyage(myvoyage);
        print('Voyage ajouté au fournisseur: $myvoyage');
        Navigator.pushNamed(context, HomeView.routeName);
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    String destinationName =
        ModalRoute.of(context)!.settings.arguments as String;
    Destination destination = Provider.of<DestinationProvider>(context)
        .getDestinationByName(destinationName);

    return Scaffold(
      appBar: AppBar(
        title: const Text('Organisation voyage'),
        actions: <Widget>[
          IconButton(
            icon: const Icon(Icons.add),
            onPressed: () => Navigator.pushNamed(
              context,
              ActivityFormView.routeName,
              arguments: destinationName,
            ),
          )
        ],
      ),
      drawer: const AppDrawer(),
      body: Column(
        children: <Widget>[
          VoyageOverview(
            destinationName: destination.name,
            voyage: myvoyage,
            setDate: setDate,
            amount: amount,
            destinationImage: destination.image,
          ),
          Expanded(
            child: index == 0
                ? ActivityList(
                    activities: destination.activities,
                    selectedActivities: myvoyage.activities,
                    toggleActivity: toggleActivity,
                  )
                : VoyageActivityList(
                    activities: myvoyage.activities,
                    deleteVoyageActivity: deleteVoyageActivity,
                  ),
          ),
        ],
      ),
      floatingActionButton: FloatingActionButton(
        child: const Icon(Icons.forward),
        onPressed: () => saveVoyage(destination.name),
      ),
      bottomNavigationBar: BottomNavigationBar(
        currentIndex: index,
        items: const [
          BottomNavigationBarItem(
            icon: Icon(Icons.map),
            label: 'Découverte',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.stars),
            label: 'Mes activités',
          )
        ],
        onTap: switchIndex,
      ),
    );
  }
}
