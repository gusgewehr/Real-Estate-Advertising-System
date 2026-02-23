import 'package:frontend/domain/entity/real_estate.dart';
import 'package:file_picker/file_picker.dart';

import '../entity/pagination.dart';

abstract class RealEstateRepository {
  Future<void> createRealEstate(RealEstateEntity realEstate);
  Future<PaginationEntity<RealEstateEntity>> getRealEstate(int page, int pageSize);
  Future<String> uploadImage(PlatformFile image);
}

