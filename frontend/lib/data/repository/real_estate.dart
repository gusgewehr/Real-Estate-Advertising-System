import 'package:frontend/domain/entity/real_estate.dart';
import 'package:frontend/domain/repository/real_estate.dart';
import 'package:file_picker/file_picker.dart';

import '../../domain/entity/pagination.dart';
import '../datasources/real_estate.dart';

class RealEstateRepositoryImpl implements RealEstateRepository{
  final RealEstateDataSource realEstateDataSource;

  RealEstateRepositoryImpl({required this.realEstateDataSource});

  @override
  Future<void> createRealEstate(RealEstateEntity realEstate) async {
    await realEstateDataSource.createRealEstate(realEstate);

  }

  @override
  Future<PaginationEntity<RealEstateEntity>> getRealEstate(int page, int pageSize) async {
    final res =  await realEstateDataSource.getRealEstate(page, pageSize);

    return res;
  }

  @override
  Future<String> uploadImage(PlatformFile image) async {
    final res = await realEstateDataSource.uploadImage(image);

    return res;
  }

}

