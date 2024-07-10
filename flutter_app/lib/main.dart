import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_app/views/groupe_detail/groupe_detail_screen.dart';
import 'package:flutter_app/views/login/login_screen.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:flutter_app/providers/group_voyage_provider.dart';
import 'package:flutter_app/views/add_group/add_group_screen.dart';
import 'package:flutter_app/views/google_map/google_map_hotel_view.dart';
import 'package:flutter_app/views/group_voyage/groupeVoyage_screen.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:flutter_app/providers/auth_provider.dart';
import 'package:flutter_app/providers/destination_provider.dart';
import 'package:flutter_app/providers/voyage_provider.dart';
import 'package:flutter_app/views/activity_fom/activity_form_view.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:flutter_app/views/profile/profile_screen.dart';
import 'package:flutter_app/views/register/register_screen.dart';
import 'package:flutter_app/views/google_map/google_map_view.dart';
import 'package:flutter_app/views/voyage/voyage_view.dart';
import 'package:provider/provider.dart';
import 'package:uni_links/uni_links.dart';
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

  static void setLocale(BuildContext context, Locale newLocale) {
    _AppVoyageState state = context.findAncestorStateOfType<_AppVoyageState>()!;
    state.setLocale(newLocale);
  }

  @override
  State<MyApp> createState() => _AppVoyageState();
}

class _AppVoyageState extends State<MyApp> {
  final AuthProvider authProvider = AuthProvider();
  final DestinationProvider destinationProvider = DestinationProvider();
  final VoyageProvider voyageProvider = VoyageProvider();
  Locale _locale = const Locale('fr', '');
  final GroupVoyageProvider groupVoyageProvider = GroupVoyageProvider();
  StreamSubscription? _sub;


  @override
  void initState() {
    voyageProvider.fetchData();
    destinationProvider.fetchData();
    super.initState();
  }


  void setLocale(Locale locale) {
    setState(() {
      _locale = locale;
    });
  }

  @override
  void dispose() {
    _sub?.cancel();
    super.dispose();
  }

  void initUniLinks() async {
    _sub = linkStream.listen((String link) async {
      // Gérer le lien profond ici
      Uri uri = Uri.parse(link);
      if (uri.pathSegments.contains('join')) {
        // Extraire les paramètres nécessaires
        int groupeId = int.parse(uri.pathSegments[1]);
        String? token = uri.queryParameters['token'];

        // Utiliser GroupProvider pour rejoindre le groupe
        try {
          await groupVoyageProvider.JoinGroup(groupeId, token);

          // Rediriger vers la page de détail du groupe
          Navigator.push(
            context,
            MaterialPageRoute(
              builder: (context) => GroupeDetailScreen(
                groupeId: groupeId,
               // token: token
              ),
            ),
          );
        } catch (e) {
          // Gérer les erreurs si la tentative de rejoindre échoue
          print('Erreur lors de la tentative de rejoindre le groupe: $e');
        }
      }
        } as void Function(String? event)?, onError: (err) {
      // Gérer les erreurs de lien profond ici
    });
  }

  @override
  Widget build(BuildContext context) {
    return MultiProvider(


      providers: [
        ChangeNotifierProvider.value(value: voyageProvider),
        ChangeNotifierProvider.value(value: destinationProvider),
        ChangeNotifierProvider.value(value: authProvider),
        ChangeNotifierProvider.value(value: groupVoyageProvider)
      ],

      child: MaterialApp(
        locale: _locale,
        localizationsDelegates: const [
          AppLocalizations.delegate,
          GlobalMaterialLocalizations.delegate,
          GlobalWidgetsLocalizations.delegate,
          GlobalCupertinoLocalizations.delegate,
        ],
        supportedLocales: const [
          Locale('fr'),
          Locale('en'),
          Locale('es'),
        ],

        debugShowCheckedModeBanner: false,
        routes: {
          '/': (context) => const LoginScreen(),
          DestinationView.routeName: (_) => const DestinationView(),
          VoyageView.routeName: (_) => const VoyageView(),
          VoyagesView.routeName: (_) => const VoyagesView(),
          ActivityFormView.routeName: (_) => const ActivityFormView(),
          ProfileScreen.routeName: (_)=> const ProfileScreen(),
          RegisterScreen.routeName: (_) => const RegisterScreen(),
          HomeView.routeName: (_) => const HomeView(),
          GoogleMapView.routeName: (_) => const GoogleMapView(),
          GoogleMapViewHotel.routeName: (_) => const GoogleMapViewHotel(),
          GroupevoyageScreen.routeName: (_) => const GroupevoyageScreen(),
          AddGroupScreen.routeName: (_) => const AddGroupScreen()
        },
        onUnknownRoute: (_) => MaterialPageRoute(
          builder: (_) => const NotFound(),
        ),
      ),
    );
  }
}
