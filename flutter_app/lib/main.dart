import 'package:flutter/material.dart';
import 'package:flutter_app/providers/destination_provider.dart';
import 'package:flutter_app/providers/voyage_provider.dart';
import 'package:flutter_app/views/activity_fom/activity_form_view.dart';
import 'package:provider/provider.dart';
import 'views/destination/destination_view.dart';
import 'views/voyages/voyages_view.dart';
import 'views/not-found/not_found.dart';
import 'views/voyage/voyage_view.dart';
import './views/home/home_view.dart';

main() => runApp(const MyApp());

class MyApp extends StatefulWidget {
  const MyApp({super.key});

  @override
  State<MyApp> createState() => _AppVoyageState();
}

class _AppVoyageState extends State<MyApp> {
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
      ],
      child: MaterialApp(
        debugShowCheckedModeBanner: false,
        routes: {
          '/': (context) => const HomeView(),
          DestinationView.routeName: (_) => const DestinationView(),
          VoyagesView.routeName: (_) => const VoyagesView(),
          VoyageView.routeName: (_) => const VoyagesView(),
          ActivityFormView.routeName: (_) => const ActivityFormView(),
        },
        onUnknownRoute: (_) => MaterialPageRoute(
          builder: (_) => const NotFound(),
        ),
      ),
    );
  }
}
