import 'package:flutter/material.dart';
import 'package:flutter_app/views/google_map/google_map_hotel_view.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:flutter_app/providers/auth_provider.dart';
import 'package:flutter_app/providers/destination_provider.dart';
import 'package:flutter_app/providers/voyage_provider.dart';
import 'package:flutter_app/views/activity_fom/activity_form_view.dart';
import 'package:flutter_app/views/login/login_screen.dart';
import 'package:flutter_app/views/profile/profile_screen.dart';
import 'package:flutter_app/views/register/register_screen.dart';
import 'package:flutter_app/views/google_map/google_map_view.dart';
import 'package:flutter_app/views/voyage/voyage_view.dart';
import 'package:provider/provider.dart';

import 'views/destination/destination_view.dart';
import 'views/voyages/voyages_view.dart';
import 'views/not-found/not_found.dart';
import './views/home/home_view.dart';



Future main() async {
  // To load the .env file contents into dotenv.
  // NOTE: fileName defaults to .env and can be omitted in this case.
  // Ensure that the filename corresponds to the path in step 1 and 2.
  await dotenv.load(fileName: ".env");
  runApp(const MyApp());
}

class MyApp extends StatefulWidget {
  const MyApp({super.key});

  @override
  State<MyApp> createState() => _AppVoyageState();
}

class _AppVoyageState extends State<MyApp> {
  final AuthProvider authProvider = AuthProvider();
  final DestinationProvider destinationProvider = DestinationProvider();
  final VoyageProvider voyageProvider = VoyageProvider();

  @override
  void initState() {
    voyageProvider.fetchData();
    destinationProvider.fetchData();
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return MultiProvider(
      providers: [
        ChangeNotifierProvider.value(value: voyageProvider),
        ChangeNotifierProvider.value(value: destinationProvider),
        ChangeNotifierProvider.value(value: authProvider),
      ],
      child: MaterialApp(
        debugShowCheckedModeBanner: false,
        routes: {
          '/': (context) => const HomeView(),
          DestinationView.routeName: (_) => const DestinationView(),
          VoyageView.routeName: (_) => const VoyageView(),
          VoyagesView.routeName: (_) => const VoyagesView(),
          ActivityFormView.routeName: (_) => const ActivityFormView(),
          ProfileScreen.routeName: (_)=> const ProfileScreen(),
          RegisterScreen.routeName: (_) => const RegisterScreen(),
          HomeView.routeName: (_) => const HomeView(),
          GoogleMapView.routeName: (_) => const GoogleMapView(),
          GoogleMapViewHotel.routeName: (_) => const GoogleMapViewHotel(),
        },
        onUnknownRoute: (_) => MaterialPageRoute(
          builder: (_) => const NotFound(),
        ),
      ),
    );
  }
}
