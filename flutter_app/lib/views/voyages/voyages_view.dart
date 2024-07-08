import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../../providers/voyage_provider.dart';
import '../../widgets/app_loader.dart';
import 'widgets/voyage_list.dart';
import '../../widgets/app_drawer.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class VoyagesView extends StatelessWidget {
  static const String routeName = '/voyages';

  const VoyagesView({super.key});

  @override
  Widget build(BuildContext context) {
    VoyageProvider voyageProvider = Provider.of<VoyageProvider>(context);
    return DefaultTabController(
      length: 2,
      child: Scaffold(
        appBar: AppBar(
          title: Text(AppLocalizations.of(context)!.mes_voyages),
          bottom: TabBar(
            tabs: <Widget>[
              Tab(
                text: AppLocalizations.of(context)!.voyage_futur,
              ),
              Tab(
                text: AppLocalizations.of(context)!.voyage_passed,
              ),
            ],
          ),
        ),
        drawer: const AppDrawer(),
        body: voyageProvider.isLoading == false
            ? voyageProvider.voyages.isNotEmpty
                ? TabBarView(
                    children: <Widget>[
                      VoyageList(
                        voyages: voyageProvider.voyages
                            .where((voyage) =>
                                DateTime.now().isBefore(voyage.dateRetour!))
                            .toList(),
                      ),
                      VoyageList(
                        voyages: voyageProvider.voyages
                            .where((voyage) =>
                                DateTime.now().isAfter(voyage.dateRetour!))
                            .toList(),
                      ),
                    ],
                  )
                : Container(
                    alignment: Alignment.center,
                    child: Text(AppLocalizations.of(context)!.aucun_voyage),
                  )
            : const AppLoader(),
      ),
    );
  }
}
