import 'package:country_flags/country_flags.dart';
import 'package:flutter/material.dart';
import 'package:flutter_app/main.dart';
import 'package:flutter_app/providers/auth_provider.dart';
import 'package:flutter_app/providers/flipping_provider.dart';
import 'package:flutter_app/views/login/login_screen.dart';
import 'package:flutter_app/views/profile/profile_screen.dart';
import 'package:provider/provider.dart';
import '../views/home/home_view.dart';
import '../views/voyages/voyages_view.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';


class AppDrawer extends StatefulWidget {
  const AppDrawer({Key? key}) : super(key: key);

  @override
  _AppDrawerState createState() => _AppDrawerState();
}

class _AppDrawerState extends State<AppDrawer> {

  Locale _locale = const Locale('fr', '');
  bool isFeatureEnabled = true;

  void _changeLanguage(Locale locale) {
    setState(() {
      _locale = locale;
    });
    MyApp.setLocale(context, locale);
  }

  void toggleFeature(bool value) async {
    try {
      await updateFeatureToggle(value);
      await Future.delayed(Duration(seconds: 2));
      if (value) {
        print("Feature activated on the server");
      } else {
        print("Feature deactivated on the server");
      }
      setState(() {
        isFeatureEnabled = value;
      });

      String message;
      if (isFeatureEnabled) {
        message = AppLocalizations.of(context)!.message_function_activate;
      } else {
        message = AppLocalizations.of(context)!.message_function_desactivate;
      }

      showDialog(
        context: context,
        builder: (context) {
          return AlertDialog(
            title: Text(AppLocalizations.of(context)!.alert_update_function),
            content: Text(message),
            actions: <Widget>[
              TextButton(
                child: Text('OK'),
                onPressed: () {
                  Navigator.of(context).pop();
                },
              ),
            ],
          );
        },
      );
    } catch (e) {
      showDialog(
        context: context,
        builder: (context) {
          return AlertDialog(
            title: Text(AppLocalizations.of(context)!.erreur),
            content: Text(AppLocalizations.of(context)!.echec_update_function),
            actions: <Widget>[
              TextButton(
                child: Text('OK'),
                onPressed: () {
                  Navigator.of(context).pop();
                },
              ),
            ],
          );
        },
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return Drawer(
      child: ListView(
        children: <Widget>[
          DrawerHeader(
            decoration: BoxDecoration(
              gradient: LinearGradient(
                colors: [
                  Theme.of(context).primaryColor,
                  Theme.of(context).primaryColor.withOpacity(0.5),
                ],
                begin: Alignment.topLeft,
                end: Alignment.bottomRight,
              ),
            ),
            child: Text(AppLocalizations.of(context)!.title_drawer,
              style: TextStyle(
                color: Colors.white,
                fontSize: 30,
              ),
            ),
          ),
          ListTile(
            leading: const Icon(Icons.home),
            title: Text(AppLocalizations.of(context)!.accueil),
            onTap: () {
              Navigator.pushNamed(context, HomeView.routeName);
            },
          ),
          ListTile(
            leading: const Icon(Icons.person),
            title: const Text('Profile'),
            onTap: () {
              Navigator.pushNamed(context, ProfileScreen.routeName);
            },
          ),
          ListTile(
            leading: const Icon(Icons.flight),
            title: Text(AppLocalizations.of(context)!.mes_voyages),
            onTap: () {
              Navigator.pushNamed(context, VoyagesView.routeName);
            },
          ),
          ListTile(
            leading: const Icon(Icons.logout),
            title: Text(AppLocalizations.of(context)!.deconnexion),
            onTap: () async {
              final authProvider = Provider.of<AuthProvider>(context, listen: false);
              await authProvider.Logout();
              Navigator.pushReplacement(
                context,
                MaterialPageRoute(builder: (BuildContext context) => const LoginScreen()),
              );
            },
          ),
          SwitchListTile(
            title: Text(AppLocalizations.of(context)!.toogle_drawer),
            value: isFeatureEnabled,
            onChanged: (bool value) {
              toggleFeature(value);
            },
          ),
          ListTile(
            leading: const Icon(Icons.language),
            title: Text(AppLocalizations.of(context)!.langage),
            trailing: DropdownButton<Locale>(
              value: _locale,
              items: [
                DropdownMenuItem(
                  value: const Locale('fr', ''),
                  child: Row(
                    children: [
                      CountryFlag.fromLanguageCode('fr',
                          width: 20,
                          height: 20),
                      const SizedBox(width: 8),
                      const Text('Français'),
                    ],
                  ),
                ),
                DropdownMenuItem(
                  value: const Locale('en', ''),
                  child: Row(
                    children: [
                      CountryFlag.fromLanguageCode('en',
                          width: 20,
                          height: 20),
                      const SizedBox(width: 8),
                      const Text('English'),
                    ],
                  ),
                ),
              ],
              onChanged: (Locale? locale) {
                if (locale != null) {
                  _changeLanguage(locale);
                }
              },
            ),
          ),
        ],
      ),
    );
  }
}