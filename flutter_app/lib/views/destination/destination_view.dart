import 'package:flutter/material.dart';
import 'package:flutter_app/models/hotel_model.dart';
import 'package:flutter_app/views/activity_fom/activity_form_view.dart';
import 'package:flutter_app/views/destination/widgets/hotel_list.dart';
import 'package:flutter_app/views/destination/widgets/voyage_activity_hotel.dart';
import 'package:flutter_app/views/destination/widgets/voyage_hotel_list.dart';
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
      hotels: [],
      dateAller: null,
      dateRetour: null,
      destination: '',
    );
  }

  double get amount {
    final activityTotal = myvoyage.activities.fold(0.0, (prev, element) {
      return prev + element.price;
    });
    final hotelTotal = myvoyage.hotels.fold(0.0, (prev, element) {
      return prev + element.price;
    });
    final totalAmount = activityTotal + hotelTotal;
    return double.parse(totalAmount.toStringAsFixed(2));
  }


  void setDateAller() {
    showDatePicker(
      context: context,
      initialDate: DateTime.now().add(const Duration(days: 1)),
      firstDate: DateTime.now(),
      lastDate: DateTime(2030),
    ).then((newDate) {
      if (newDate != null) {
        setState(() {
          myvoyage.dateAller = newDate;
          myvoyage.dateRetour = null; // Reset the return date if the departure date is changed
        });
      }
    });
  }

  void setDateRetour() {
    if (myvoyage.dateAller == null) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          content: Text('Veuillez d\'abord sélectionner une date aller'),
          backgroundColor: Colors.red,
        ),
      );
      return;
    }

    showDatePicker(
      context: context,
      initialDate: myvoyage.dateAller!.add(const Duration(days: 1)),
      firstDate: myvoyage.dateAller!.add(const Duration(days: 1)),
      lastDate: DateTime(2030),
    ).then((newDate) {
      if (newDate != null) {
        setState(() {
          myvoyage.dateRetour = newDate;
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


  void toggleHotel(Hotel hotel) {
    setState(() {
      myvoyage.hotels.contains(hotel)
          ? myvoyage.hotels.remove(hotel)
          : myvoyage.hotels.add(hotel);
    });
  }

  void deleteVoyageHotel(Hotel hotel) {
    setState(() {
      myvoyage.hotels.remove(hotel);
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

    if (myvoyage.dateAller == null || myvoyage.dateRetour == null) {
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
    final routeArgs = ModalRoute.of(context)?.settings.arguments;
    if (routeArgs == null) {
      return Scaffold(
        appBar: AppBar(
          title: const Text('Erreur'),
        ),
        body: const Center(
          child: Text('Aucune destination fournie.'),
        ),
      );
    }

    String destinationName = routeArgs as String;
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
            destinationImage: destination.image,
            voyage: myvoyage,
            setDateAller: setDateAller,
            setDateRetour: setDateRetour,
            amount: amount,
          ),
          Expanded(
            child: index == 0
                ? VoyageActivityHotelList(
              activities: destination.activities,
              selectedActivities: myvoyage.activities,
              toggleActivity: toggleActivity,
              hotels: destination.hotels,
              selectedHotels: myvoyage.hotels,
              toggleHotel: toggleHotel,
            )
                : index == 1
                ? VoyageActivityList(
              activities: myvoyage.activities,
              deleteVoyageActivity: deleteVoyageActivity,

            )
                : index == 2
                ? VoyageHotelList(
              hotels: myvoyage.hotels,
              deleteVoyageHotel: deleteVoyageHotel,
            )
            :
            Container(),
          ),
        ],
      ),
      floatingActionButton: FloatingActionButton(
        child: const Icon(Icons.forward),
        onPressed: () => saveVoyage(destinationName),
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
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.hotel),
            label: 'Mes hôtels',
          ),
        ],
        onTap: switchIndex,
      ),
    );
  }
}






