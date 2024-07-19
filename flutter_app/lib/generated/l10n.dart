// GENERATED CODE - DO NOT MODIFY BY HAND
import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'intl/messages_all.dart';

// **************************************************************************
// Generator: Flutter Intl IDE plugin
// Made by Localizely
// **************************************************************************

// ignore_for_file: non_constant_identifier_names, lines_longer_than_80_chars
// ignore_for_file: join_return_with_assignment, prefer_final_in_for_each
// ignore_for_file: avoid_redundant_argument_values, avoid_escaping_inner_quotes

class S {
  S();

  static S? _current;

  static S get current {
    assert(_current != null,
        'No instance of S was loaded. Try to initialize the S delegate before accessing S.current.');
    return _current!;
  }

  static const AppLocalizationDelegate delegate = AppLocalizationDelegate();

  static Future<S> load(Locale locale) {
    final name = (locale.countryCode?.isEmpty ?? false)
        ? locale.languageCode
        : locale.toString();
    final localeName = Intl.canonicalizedLocale(name);
    return initializeMessages(localeName).then((_) {
      Intl.defaultLocale = localeName;
      final instance = S();
      S._current = instance;

      return instance;
    });
  }

  static S of(BuildContext context) {
    final instance = S.maybeOf(context);
    assert(instance != null,
        'No instance of S present in the widget tree. Did you add S.delegate in localizationsDelegates?');
    return instance!;
  }

  static S? maybeOf(BuildContext context) {
    return Localizations.of<S>(context, S);
  }

  /// `TRIP ESGI`
  String get title_drawer {
    return Intl.message(
      'TRIP ESGI',
      name: 'title_drawer',
      desc: '',
      args: [],
    );
  }

  /// `Home`
  String get accueil {
    return Intl.message(
      'Home',
      name: 'accueil',
      desc: '',
      args: [],
    );
  }

  /// `My trips`
  String get mes_voyages {
    return Intl.message(
      'My trips',
      name: 'mes_voyages',
      desc: '',
      args: [],
    );
  }

  /// `Profile`
  String get profile {
    return Intl.message(
      'Profile',
      name: 'profile',
      desc: '',
      args: [],
    );
  }

  /// `Logout`
  String get deconnexion {
    return Intl.message(
      'Logout',
      name: 'deconnexion',
      desc: '',
      args: [],
    );
  }

  /// `Activate/Disable Trip`
  String get toogle_drawer {
    return Intl.message(
      'Activate/Disable Trip',
      name: 'toogle_drawer',
      desc: '',
      args: [],
    );
  }

  /// `Langage`
  String get langage {
    return Intl.message(
      'Langage',
      name: 'langage',
      desc: '',
      args: [],
    );
  }

  /// `Trip save feature is disabled`
  String get alert_function_flipping {
    return Intl.message(
      'Trip save feature is disabled',
      name: 'alert_function_flipping',
      desc: '',
      args: [],
    );
  }

  /// `Do you want to save?`
  String get sauvegarde_voyage {
    return Intl.message(
      'Do you want to save?',
      name: 'sauvegarde_voyage',
      desc: '',
      args: [],
    );
  }

  /// `Cancem`
  String get annulation_voyage {
    return Intl.message(
      'Cancem',
      name: 'annulation_voyage',
      desc: '',
      args: [],
    );
  }

  /// `Save`
  String get sauvegarde {
    return Intl.message(
      'Save',
      name: 'sauvegarde',
      desc: '',
      args: [],
    );
  }

  /// `Wait !`
  String get attention {
    return Intl.message(
      'Wait !',
      name: 'attention',
      desc: '',
      args: [],
    );
  }

  /// `Please select a departure date first`
  String get alert_date {
    return Intl.message(
      'Please select a departure date first',
      name: 'alert_date',
      desc: '',
      args: [],
    );
  }

  /// `You have not entered a date`
  String get attention_content {
    return Intl.message(
      'You have not entered a date',
      name: 'attention_content',
      desc: '',
      args: [],
    );
  }

  /// `Error saving trip`
  String get erreur_sauvegarde {
    return Intl.message(
      'Error saving trip',
      name: 'erreur_sauvegarde',
      desc: '',
      args: [],
    );
  }

  /// `Error`
  String get erreur {
    return Intl.message(
      'Error',
      name: 'erreur',
      desc: '',
      args: [],
    );
  }

  /// `No destination provided.`
  String get empty_destination {
    return Intl.message(
      'No destination provided.',
      name: 'empty_destination',
      desc: '',
      args: [],
    );
  }

  /// `Travel organization`
  String get organisation_titre {
    return Intl.message(
      'Travel organization',
      name: 'organisation_titre',
      desc: '',
      args: [],
    );
  }

  /// `Discovery`
  String get label_decouverte {
    return Intl.message(
      'Discovery',
      name: 'label_decouverte',
      desc: '',
      args: [],
    );
  }

  /// `My activities`
  String get label_activites {
    return Intl.message(
      'My activities',
      name: 'label_activites',
      desc: '',
      args: [],
    );
  }

  /// `My hôtels`
  String get label_hotels {
    return Intl.message(
      'My hôtels',
      name: 'label_hotels',
      desc: '',
      args: [],
    );
  }

  /// `Feature enabled`
  String get message_function_activate {
    return Intl.message(
      'Feature enabled',
      name: 'message_function_activate',
      desc: '',
      args: [],
    );
  }

  /// `Feature Update`
  String get message_function_desactivate {
    return Intl.message(
      'Feature Update',
      name: 'message_function_desactivate',
      desc: '',
      args: [],
    );
  }

  /// `Update feature`
  String get alert_update_function {
    return Intl.message(
      'Update feature',
      name: 'alert_update_function',
      desc: '',
      args: [],
    );
  }

  /// `Failed to update feature`
  String get echec_update_function {
    return Intl.message(
      'Failed to update feature',
      name: 'echec_update_function',
      desc: '',
      args: [],
    );
  }

  /// `List of activities`
  String get title_activities {
    return Intl.message(
      'List of activities',
      name: 'title_activities',
      desc: '',
      args: [],
    );
  }

  /// `List of Hotels`
  String get title_hotels {
    return Intl.message(
      'List of Hotels',
      name: 'title_hotels',
      desc: '',
      args: [],
    );
  }

  /// `Activity deleted`
  String get delete_activite {
    return Intl.message(
      'Activity deleted',
      name: 'delete_activite',
      desc: '',
      args: [],
    );
  }

  /// `Hotel deleted`
  String get delete_hotel {
    return Intl.message(
      'Hotel deleted',
      name: 'delete_hotel',
      desc: '',
      args: [],
    );
  }

  /// `Date to Go`
  String get date_aller {
    return Intl.message(
      'Date to Go',
      name: 'date_aller',
      desc: '',
      args: [],
    );
  }

  /// `Return Date`
  String get date_retour {
    return Intl.message(
      'Return Date',
      name: 'date_retour',
      desc: '',
      args: [],
    );
  }

  /// `Select a date`
  String get selected_date {
    return Intl.message(
      'Select a date',
      name: 'selected_date',
      desc: '',
      args: [],
    );
  }

  /// `Amount/person`
  String get montant_personne {
    return Intl.message(
      'Amount/person',
      name: 'montant_personne',
      desc: '',
      args: [],
    );
  }

  /// `Esgi-Trip`
  String get title_home {
    return Intl.message(
      'Esgi-Trip',
      name: 'title_home',
      desc: '',
      args: [],
    );
  }

  /// `Search for a city`
  String get recherche_ville {
    return Intl.message(
      'Search for a city',
      name: 'recherche_ville',
      desc: '',
      args: [],
    );
  }

  /// `No result`
  String get resultat {
    return Intl.message(
      'No result',
      name: 'resultat',
      desc: '',
      args: [],
    );
  }

  /// `No journey is over yet`
  String get statut_voyage {
    return Intl.message(
      'No journey is over yet',
      name: 'statut_voyage',
      desc: '',
      args: [],
    );
  }

  /// `In progress`
  String get statut_in_progress {
    return Intl.message(
      'In progress',
      name: 'statut_in_progress',
      desc: '',
      args: [],
    );
  }

  /// `Finished`
  String get statut_terminer {
    return Intl.message(
      'Finished',
      name: 'statut_terminer',
      desc: '',
      args: [],
    );
  }

  /// `Missing or invalid arguments`
  String get args_manquants_invalides {
    return Intl.message(
      'Missing or invalid arguments',
      name: 'args_manquants_invalides',
      desc: '',
      args: [],
    );
  }

  /// `Invalid arguments`
  String get args_invalides {
    return Intl.message(
      'Invalid arguments',
      name: 'args_invalides',
      desc: '',
      args: [],
    );
  }

  /// `DestinationProvider not found`
  String get destination_prov_not_found {
    return Intl.message(
      'DestinationProvider not found',
      name: 'destination_prov_not_found',
      desc: '',
      args: [],
    );
  }

  /// `Destination not found`
  String get destination_not_found {
    return Intl.message(
      'Destination not found',
      name: 'destination_not_found',
      desc: '',
      args: [],
    );
  }

  /// `Future`
  String get voyage_futur {
    return Intl.message(
      'Future',
      name: 'voyage_futur',
      desc: '',
      args: [],
    );
  }

  /// `Past`
  String get voyage_passed {
    return Intl.message(
      'Past',
      name: 'voyage_passed',
      desc: '',
      args: [],
    );
  }

  /// `No travel at the moment`
  String get aucun_voyage {
    return Intl.message(
      'No travel at the moment',
      name: 'aucun_voyage',
      desc: '',
      args: [],
    );
  }
}

class AppLocalizationDelegate extends LocalizationsDelegate<S> {
  const AppLocalizationDelegate();

  List<Locale> get supportedLocales {
    return const <Locale>[
      Locale.fromSubtags(languageCode: 'en'),
    ];
  }

  @override
  bool isSupported(Locale locale) => _isSupported(locale);
  @override
  Future<S> load(Locale locale) => S.load(locale);
  @override
  bool shouldReload(AppLocalizationDelegate old) => false;

  bool _isSupported(Locale locale) {
    for (var supportedLocale in supportedLocales) {
      if (supportedLocale.languageCode == locale.languageCode) {
        return true;
      }
    }
    return false;
  }
}
